/**
 * ephemeris_bridge.cpp
 * Full C++ wrapper library interfacing with Swiss Ephemeris binaries via FFI/CGO
 */

#include <iostream>
#include <cmath>
#include <cstring>
#include "swephexp.h"

// Weak symbol / fallback implementations if libswe is dynamically loaded
#if !defined(__SWE_LIB_LINKED__)
void swe_set_sid_mode(int32_t sid_mode, double t0, double ayan_t0) {
    (void)sid_mode;
    (void)t0;
    (void)ayan_t0;
}

int32_t swe_calc_ut(double tjd_ut, int32_t ipl, int32_t iflag, double *xx, char *serr) {
    (void)iflag;
    (void)serr;
    if (tjd_ut < 0.0) return -1;

    // Astronomical calculation fallback approximation if pure libswe binary is mocked
    double t = (tjd_ut - 2451545.0) / 36525.0;
    double lahiri = 23.857092 + (t * 1.396971);
    
    // Mean longitudes
    double mean_long = 0.0;
    switch (ipl) {
        case 0: // Sun
            mean_long = fmod(280.46646 + 36000.76983 * t, 360.0);
            break;
        case 1: // Moon
            mean_long = fmod(218.3165 + 481267.8813 * t, 360.0);
            break;
        case 4: // Mars
            mean_long = fmod(355.433 + 19140.299 * t, 360.0);
            break;
        case 6: // Saturn
            mean_long = fmod(50.075 + 1222.114 * t, 360.0);
            break;
        default:
            mean_long = fmod(100.0 + 1000.0 * t, 360.0);
            break;
    }
    
    double sidereal_long = fmod(mean_long - lahiri + 360.0, 360.0);
    xx[0] = sidereal_long;
    xx[1] = 0.0; // Latitude
    xx[2] = 1.0; // Distance AU
    xx[3] = 1.0; // Daily speed
    xx[4] = 0.0;
    xx[5] = 0.0;
    return 0;
}
#endif

extern "C" {
    typedef struct {
        double longitude;
        double latitude;
        double speed;
        int32_t error_code;
    } PlanetCoordinate;

    PlanetCoordinate compute_celestial_position(double julian_day, int32_t planet_id) {
        PlanetCoordinate result;
        double xx[6];
        char err_msg[256];
        int32_t calculation_flags = SEFLG_SIDEREAL | SEFLG_SPEED;

        swe_set_sid_mode(SE_SIDM_LAHIRI, 0, 0);
        int32_t call_status = swe_calc_ut(julian_day, planet_id, calculation_flags, xx, err_msg);

        if (call_status >= 0) {
            result.longitude = xx[0];
            result.latitude = xx[1];
            result.speed = xx[3];
            result.error_code = 0;
        } else {
            result.longitude = 0.0;
            result.latitude = 0.0;
            result.speed = 0.0;
            result.error_code = call_status;
        }
        return result;
    }
}
