/**
 * Conflict Resolver Matrix Modal
 * Detailed mathematical visualization of how the Symbolic Reasoning Engine
 * resolves conflicting yogas into a single Net Auspiciousness Vector (A_net).
 */

import React from 'react';
import { ConflictResolutionResult } from '../engine/conflictResolver';

interface ConflictResolverModalProps {
  isOpen: boolean;
  onClose: () => void;
  resolution: ConflictResolutionResult;
}

export const ConflictResolverModal: React.FC<ConflictResolverModalProps> = ({
  isOpen,
  onClose,
  resolution,
}) => {
  if (!isOpen) return null;

  const beneficYogas = resolution.allTriggeredYogas.filter(y => y.nature === 'Benefic');
  const maleficYogas = resolution.allTriggeredYogas.filter(y => y.nature === 'Malefic');

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-[#0e0205]/90 backdrop-blur-md animate-fade-in">
      <div className="relative w-full max-w-3xl max-h-[90vh] overflow-y-auto rounded-xl border border-rose-900/60 bg-[#1e050b] p-6 shadow-2xl text-rose-100">
        {/* Header */}
        <div className="flex items-start justify-between border-b border-rose-950/70 pb-4">
          <div>
            <h2 className="font-serif text-xl font-bold text-amber-300">
              Symbolic Conflict Resolver Matrix
            </h2>
            <p className="mt-1 text-xs text-rose-300/60">
              Eliminating contradictory logic loops through algebraic vector subtraction & dasha modulation
            </p>
          </div>
          <button
            onClick={onClose}
            className="rounded-lg p-1.5 text-rose-300/60 hover:bg-[#2e0914] hover:text-white transition-colors"
          >
            ✕
          </button>
        </div>

        {/* Content */}
        <div className="mt-5 space-y-6">
          {/* Step 1: Opposing Forces Diagram */}
          <div className="rounded-lg border border-rose-950/70 bg-[#25070f] p-4">
            <div className="text-xs font-semibold text-rose-100">
              Step 1: Raw Potency Evaluation (P_net)
            </div>
            <div className="mt-1 text-xs text-rose-300/60 font-mono">
              P_net = Σ(Positive Yogas × Strength) - Σ(Negative Yogas × Strength)
            </div>

            <div className="mt-4 grid grid-cols-1 md:grid-cols-2 gap-4">
              {/* Benefic vector */}
              <div className="rounded-lg border border-emerald-900/40 bg-emerald-950/20 p-3">
                <div className="flex items-center justify-between">
                  <span className="text-xs font-semibold text-emerald-400">Positive Yogas (Expansion)</span>
                  <span className="font-mono text-sm font-bold text-emerald-400">
                    +{beneficYogas.reduce((a, b) => a + b.calculatedStrength, 0).toFixed(1)}
                  </span>
                </div>
                <div className="mt-2 space-y-2">
                  {beneficYogas.map(y => (
                    <div key={y.name} className="text-xs border-t border-emerald-900/30 pt-1 text-rose-100/90">
                      <div className="flex justify-between font-mono">
                        <span>{y.name}</span>
                        <span>{y.calculatedStrength}</span>
                      </div>
                      <div className="text-[11px] text-rose-300/60 mt-0.5">{y.details}</div>
                    </div>
                  ))}
                </div>
              </div>

              {/* Malefic vector */}
              <div className="rounded-lg border border-rose-900/40 bg-rose-950/30 p-3">
                <div className="flex items-center justify-between">
                  <span className="text-xs font-semibold text-rose-400">Friction Yogas (Constraint)</span>
                  <span className="font-mono text-sm font-bold text-rose-400">
                    -{maleficYogas.reduce((a, b) => a + b.calculatedStrength, 0).toFixed(1)}
                  </span>
                </div>
                <div className="mt-2 space-y-2">
                  {maleficYogas.map(y => (
                    <div key={y.name} className="text-xs border-t border-rose-900/30 pt-1 text-rose-100/90">
                      <div className="flex justify-between font-mono">
                        <span>{y.name}</span>
                        <span>{y.calculatedStrength}</span>
                      </div>
                      <div className="text-[11px] text-rose-300/60 mt-0.5">{y.details}</div>
                    </div>
                  ))}
                  {maleficYogas.length === 0 && (
                    <div className="text-xs text-rose-300/50 italic">No primary friction yogas triggered</div>
                  )}
                </div>
              </div>
            </div>

            <div className="mt-3 flex items-center justify-between rounded bg-[#18050a] px-3 py-2 text-xs border border-rose-950/60">
              <span className="text-rose-200">Calculated Raw Net Potency (P_net):</span>
              <span className="font-mono text-sm font-bold text-amber-300">
                {resolution.netPotencyScore >= 0 ? `+${resolution.netPotencyScore}` : resolution.netPotencyScore}
              </span>
            </div>
          </div>

          {/* Step 2: Dasha & Transit Modulation */}
          <div className="rounded-lg border border-rose-950/70 bg-[#25070f] p-4">
            <div className="text-xs font-semibold text-rose-100">
              Step 2: Time-Varying Filter (Dasha-Transit Modulator)
            </div>
            <div className="mt-1 text-xs text-rose-300/60 font-mono">
              A_net = Clamp( (P_net / 100.0) × (M_dasha + M_transit), 0.0, 1.0 )
            </div>

            <div className="mt-3 grid grid-cols-2 md:grid-cols-4 gap-3 text-xs">
              <div className="rounded border border-rose-950/70 bg-[#1c050b] p-2.5">
                <div className="text-rose-300/60">Active Mahadasha</div>
                <div className="font-mono font-semibold text-amber-300 mt-0.5">
                  {resolution.dashaTransitModulators.activeMahadashaLord}
                </div>
              </div>
              <div className="rounded border border-rose-950/70 bg-[#1c050b] p-2.5">
                <div className="text-rose-300/60">Dasha Modulator (M_d)</div>
                <div className="font-mono font-semibold text-rose-100 mt-0.5">
                  {resolution.dashaTransitModulators.mDasha}
                </div>
              </div>
              <div className="rounded border border-rose-950/70 bg-[#1c050b] p-2.5">
                <div className="text-rose-300/60">Transit Modulator (M_t)</div>
                <div className="font-mono font-semibold text-rose-100 mt-0.5">
                  {resolution.dashaTransitModulators.mTransit >= 0 ? `+${resolution.dashaTransitModulators.mTransit}` : resolution.dashaTransitModulators.mTransit}
                </div>
              </div>
              <div className="rounded border border-amber-900/40 bg-amber-950/30 p-2.5">
                <div className="text-amber-400 font-medium">Final A_net Vector</div>
                <div className="font-mono font-bold text-amber-300 mt-0.5 text-sm">
                  {(resolution.netAuspiciousnessVector * 100).toFixed(2)}%
                </div>
              </div>
            </div>
          </div>

          {/* Step 3: Graph Traversal & Synthetic JSON Payload */}
          <div className="rounded-lg border border-rose-950/70 bg-[#25070f] p-4">
            <div className="flex items-center justify-between">
              <div className="text-xs font-semibold text-rose-100">
                Step 3: Clean Vector Payload for Desha-Kala-Patra AI Layer
              </div>
              <div className="font-mono text-[11px] text-rose-300/60">
                Latency: {resolution.graphTraversalStats.executionLatencyMs}ms · {resolution.graphTraversalStats.yogaEdgesTraversed} edges
              </div>
            </div>

            <pre className="mt-3 overflow-x-auto rounded bg-[#140407] p-3 font-mono text-[11px] text-amber-200/90 leading-relaxed border border-rose-950/60">
{JSON.stringify({
  metric: resolution.metric,
  net_auspiciousness_score: resolution.netAuspiciousnessVector,
  conflict_resolved: resolution.conflictResolved,
  dominant_vector: resolution.dominantVector,
  suppressed_vector: resolution.suppressedVector,
  mathematical_context: resolution.mathematicalContext
}, null, 2)}
            </pre>
          </div>
        </div>

        {/* Footer */}
        <div className="mt-6 flex justify-end">
          <button
            onClick={onClose}
            className="rounded-lg bg-[#2e0914] border border-rose-900/50 px-4 py-2 text-xs font-semibold text-rose-200 hover:bg-[#3d0d1b] transition-colors"
          >
            Close Resolution Details
          </button>
        </div>
      </div>
    </div>
  );
};
