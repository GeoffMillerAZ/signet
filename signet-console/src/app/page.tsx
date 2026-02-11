'use client'

import React, { useMemo } from 'react';
import ReactFlow, { 
  Background, 
  Controls, 
  Edge, 
  Node,
  Handle,
  Position
} from 'reactflow';
import 'reactflow/dist/style.css';
import { ShieldCheckIcon, CubeIcon, CodeBracketIcon, ClipboardDocumentIcon } from '@heroicons/react/24/outline';

const stats = [
  { name: 'Sealed Artifacts', stat: '12', icon: ShieldCheckIcon },
  { name: 'Active Policies', stat: '4', icon: ClipboardDocumentIcon },
  { name: 'Governance Checks', stat: '142', icon: CodeBracketIcon },
]

// Custom Node component for a "cool" look
const CustomNode = ({ data }: any) => (
  <div className="px-4 py-2 shadow-lg rounded-md bg-gray-800 border border-indigo-500/50 min-w-[150px]">
    <Handle type="target" position={Position.Top} className="w-3 h-3 bg-indigo-500" />
    <div className="flex items-center">
      <div className="rounded-full bg-indigo-500/20 p-2">
        <data.icon className="h-5 w-5 text-indigo-400" />
      </div>
      <div className="ml-3">
        <div className="text-xs font-bold text-gray-400 uppercase tracking-tighter">{data.label}</div>
        <div className="text-sm font-semibold text-white truncate max-w-[100px]">{data.value}</div>
      </div>
    </div>
    <Handle type="source" position={Position.Bottom} className="w-3 h-3 bg-indigo-500" />
  </div>
);

const nodeTypes = {
  custom: CustomNode,
};

const initialNodes: Node[] = [
  { 
    id: 'root', 
    type: 'custom',
    position: { x: 250, y: 0 }, 
    data: { label: 'Merkle Root', value: 'sha256:e3b0c44...', icon: ShieldCheckIcon } 
  },
  { 
    id: 'g1', 
    type: 'custom',
    position: { x: 50, y: 150 }, 
    data: { label: 'Persistence', value: 'sha256:f1a2b3c...', icon: CubeIcon } 
  },
  { 
    id: 'g2', 
    type: 'custom',
    position: { x: 450, y: 150 }, 
    data: { label: 'Networking', value: 'sha256:d4e5f6g...', icon: CodeBracketIcon } 
  },
  { 
    id: 'e1', 
    type: 'custom',
    position: { x: -50, y: 300 }, 
    data: { label: 'DB Schema', value: 'PASS', icon: ClipboardDocumentIcon } 
  },
  { 
    id: 'e2', 
    type: 'custom',
    position: { x: 150, y: 300 }, 
    data: { label: 'Models', value: 'WARN', icon: ClipboardDocumentIcon } 
  },
];

const initialEdges: Edge[] = [
  { id: 'e-root-g1', source: 'root', target: 'g1', animated: true },
  { id: 'e-root-g2', source: 'root', target: 'g2', animated: true },
  { id: 'e-g1-e1', source: 'g1', target: 'e1' },
  { id: 'e-g1-e2', source: 'g1', target: 'e2' },
];

export default function Home() {
  return (
    <div className="space-y-8">
      {/* Header Section */}
      <div className="md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl font-bold leading-7 text-white sm:truncate sm:text-3xl sm:tracking-tight">
            Governance Dashboard
          </h2>
          <p className="mt-1 text-sm text-gray-400">
            Real-time status of cryptographic seals and artifact integrity.
          </p>
        </div>
        <div className="mt-4 flex md:ml-4 md:mt-0">
          <button
            type="button"
            className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            New Promotion
          </button>
        </div>
      </div>

      {/* Stats Section */}
      <dl className="grid grid-cols-1 gap-5 sm:grid-cols-3">
        {stats.map((item) => (
          <div
            key={item.name}
            className="overflow-hidden rounded-lg bg-gray-900 px-4 py-5 shadow-sm border border-white/5 sm:p-6"
          >
            <div className="flex items-center">
                <item.icon className="h-6 w-6 text-indigo-400" />
                <dt className="ml-3 truncate text-sm font-medium text-gray-400">{item.name}</dt>
            </div>
            <dd className="mt-1 text-3xl font-semibold tracking-tight text-white">{item.stat}</dd>
          </div>
        ))}
      </dl>

      {/* Merkle visualization */}
      <div className="rounded-xl bg-gray-900 border border-white/5 overflow-hidden">
        <div className="px-6 py-4 border-b border-white/5 bg-gray-900/50 flex justify-between items-center">
            <h3 className="text-sm font-semibold text-white uppercase tracking-wider">Latest Seal Integrity Tree</h3>
            <span className="text-xs text-gray-500 font-mono">SEAL: 4a2b3c4d-5e6f-7g8h</span>
        </div>
        <div style={{ height: '500px' }} className="bg-gray-950/50">
          <ReactFlow
            nodes={initialNodes}
            edges={initialEdges}
            nodeTypes={nodeTypes}
            fitView
          >
            <Background color="#333" gap={20} />
            <Controls />
          </ReactFlow>
        </div>
      </div>
    </div>
  );
}