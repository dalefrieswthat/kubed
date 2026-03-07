---
layout: default
title: Home
nav_order: 1
description: "CLI and context tooling for Kubernetes, Docker, Terraform, and Helm — agent-friendly infra layout index and shell productivity"
permalink: /
---

<section class="mb-16">
  <h1 class="text-4xl md:text-5xl font-semibold text-zinc-900 tracking-tight mb-4">
    Kubed
  </h1>
  <p class="text-xl text-zinc-600 max-w-2xl leading-relaxed">
    CLI and context tooling for Kubernetes, Docker, Terraform, and Helm. A file-based, agent-friendly infra layout index so tools don't need to run <code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-base">kubectl</code> for discovery—plus shell completions and aliases to keep you in flow.
  </p>
</section>

<section class="mb-16">
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
    <div class="bg-white rounded-2xl border border-zinc-200 p-6 text-center shadow-card">
      <div class="text-4xl font-bold text-accent mb-1">1 file</div>
      <div class="text-sm text-zinc-600">instead of 10+ discovery calls</div>
    </div>
    <div class="bg-white rounded-2xl border border-zinc-200 p-6 text-center shadow-card">
      <div class="text-4xl font-bold text-accent mb-1">~1.5K tokens</div>
      <div class="text-sm text-zinc-600">layout snapshot</div>
    </div>
    <div class="bg-white rounded-2xl border border-zinc-200 p-6 text-center shadow-card">
      <div class="text-4xl font-bold text-accent mb-1">vs 50K+</div>
      <div class="text-sm text-zinc-600">typical discovery output</div>
    </div>
  </div>

  <h2 class="text-sm font-semibold text-zinc-600 uppercase tracking-wider mb-4">The problem</h2>
  <p class="text-zinc-600 max-w-2xl mb-4">
    AI agents burn thousands of tokens just to understand your infrastructure. Every <code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">kubectl get</code>, <code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">find</code>, and <code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">ls -laR</code> adds latency and cost. Discovery-heavy workflows can push context into the tens of thousands of tokens before the agent answers one question.
  </p>

  <h2 class="text-sm font-semibold text-zinc-600 uppercase tracking-wider mb-4 mt-8">The solution: two small files</h2>
  
  <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
    <div class="bg-zinc-50 rounded-xl border border-zinc-200 p-5">
      <h3 class="font-semibold text-zinc-900 mb-2"><code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">.kubed/layout.json</code></h3>
      <p class="text-sm text-zinc-600 mb-2">Static snapshot of your infrastructure:</p>
      <ul class="text-sm text-zinc-600 list-disc list-inside space-y-1">
        <li>Dockerfiles, Terraform, Helm charts</li>
        <li>Project structure</li>
        <li>Cross-repo shared infra</li>
        <li>K8s resources and relationships</li>
      </ul>
      <p class="text-xs text-zinc-500 mt-3">Run <code class="px-1 py-0.5 rounded bg-zinc-200 font-mono">kubed layout capture</code> to generate</p>
    </div>
    <div class="bg-zinc-50 rounded-xl border border-zinc-200 p-5">
      <h3 class="font-semibold text-zinc-900 mb-2"><code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">.kubed/learned.json</code></h3>
      <p class="text-sm text-zinc-600 mb-2">Accumulated knowledge across sessions:</p>
      <ul class="text-sm text-zinc-600 list-disc list-inside space-y-1">
        <li>Architecture facts ("API uses PostgreSQL")</li>
        <li>Important paths with descriptions</li>
        <li>Tech stack and dependencies</li>
        <li>Code patterns and conventions</li>
      </ul>
      <p class="text-xs text-zinc-500 mt-3">Run <code class="px-1 py-0.5 rounded bg-zinc-200 font-mono">kubed learned add-fact "..."</code> to persist</p>
    </div>
  </div>

  <p class="text-zinc-600 max-w-2xl mb-4">
    Both files are section-based with IDs and tags, so agents can query specific parts without loading everything. When the agent discovers something useful, it persists it to <code class="px-1.5 py-0.5 rounded bg-zinc-200 font-mono text-sm">learned.json</code> so the next session doesn't re-discover.
  </p>
  <p class="text-zinc-600 max-w-2xl">
    For humans, Kubed installs shell completions and aliases for Docker, Kubernetes, Terraform, and Helm so you stay productive at the terminal.
  </p>
</section>

<section class="mb-16">
  <h2 class="text-sm font-semibold text-zinc-600 uppercase tracking-wider mb-4">How it reduces tokens</h2>
  <div class="overflow-x-auto">
    <table class="w-full text-sm border border-zinc-200 rounded-lg overflow-hidden">
      <thead class="bg-zinc-100">
        <tr>
          <th class="text-left px-4 py-3 font-semibold text-zinc-700">Without Kubed</th>
          <th class="text-left px-4 py-3 font-semibold text-zinc-700">With Kubed</th>
        </tr>
      </thead>
      <tbody class="bg-white">
        <tr class="border-t border-zinc-100">
          <td class="px-4 py-3 text-zinc-600">Agent runs <code class="px-1 py-0.5 rounded bg-zinc-100 font-mono text-xs">find</code>, <code class="px-1 py-0.5 rounded bg-zinc-100 font-mono text-xs">ls -laR</code>, <code class="px-1 py-0.5 rounded bg-zinc-100 font-mono text-xs">kubectl get</code></td>
          <td class="px-4 py-3 text-zinc-600">Agent reads <code class="px-1 py-0.5 rounded bg-zinc-100 font-mono text-xs">layout.json</code></td>
        </tr>
        <tr class="border-t border-zinc-100">
          <td class="px-4 py-3 text-zinc-600">Discovery output: 50,000+ tokens</td>
          <td class="px-4 py-3 text-zinc-600">Layout file: ~1,500–3,000 tokens</td>
        </tr>
        <tr class="border-t border-zinc-100">
          <td class="px-4 py-3 text-zinc-600">Re-discovers same facts every session</td>
          <td class="px-4 py-3 text-zinc-600">Reads <code class="px-1 py-0.5 rounded bg-zinc-100 font-mono text-xs">learned.json</code> (~500 tokens)</td>
        </tr>
        <tr class="border-t border-zinc-100">
          <td class="px-4 py-3 text-zinc-600">10+ tool calls per question</td>
          <td class="px-4 py-3 text-zinc-600">1–2 file reads</td>
        </tr>
      </tbody>
    </table>
  </div>
</section>

<section class="mb-16">
  <h2 class="text-sm font-semibold text-zinc-600 uppercase tracking-wider mb-6">Tools</h2>
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
    <a href="/docker" class="group block bg-white rounded-2xl border border-zinc-200 p-6 shadow-card hover:shadow-soft hover:border-accent/40 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#2496ED]/10 flex items-center justify-center mb-4 group-hover:bg-[#2496ED]/20 transition-colors">
        <i class="fab fa-docker text-[#2496ED] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-zinc-900 mb-1">Docker</h3>
      <p class="text-sm text-zinc-600">Container management and orchestration</p>
    </a>
    <a href="/kubernetes" class="group block bg-white rounded-2xl border border-zinc-200 p-6 shadow-card hover:shadow-soft hover:border-accent/40 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#326CE5]/10 flex items-center justify-center mb-4 group-hover:bg-[#326CE5]/20 transition-colors">
        <i class="fas fa-dharmachakra text-[#326CE5] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-zinc-900 mb-1">Kubernetes</h3>
      <p class="text-sm text-zinc-600">Container orchestration and management</p>
    </a>
    <a href="/terraform" class="group block bg-white rounded-2xl border border-zinc-200 p-6 shadow-card hover:shadow-soft hover:border-accent/40 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#7B42BC]/10 flex items-center justify-center mb-4 group-hover:bg-[#7B42BC]/20 transition-colors">
        <i class="fas fa-cube text-[#7B42BC] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-zinc-900 mb-1">Terraform</h3>
      <p class="text-sm text-zinc-600">Infrastructure as Code</p>
    </a>
    <a href="/helm" class="group block bg-white rounded-2xl border border-zinc-200 p-6 shadow-card hover:shadow-soft hover:border-accent/40 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#0F1689]/10 flex items-center justify-center mb-4 group-hover:bg-[#0F1689]/20 transition-colors">
        <i class="fas fa-chart-pie text-[#0F1689] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-zinc-900 mb-1">Helm</h3>
      <p class="text-sm text-zinc-600">Kubernetes package management</p>
    </a>
  </div>
</section>

<section class="bg-white rounded-2xl border border-zinc-200 p-8 md:p-10 shadow-card mb-12">
  <h2 class="text-2xl font-semibold text-zinc-900 mb-2">Get started</h2>
  <p class="text-zinc-600 mb-6">Install Kubed and run setup for your shell. Restart your terminal after installation for changes to take effect.</p>
  <div class="space-y-4">
    <div class="rounded-xl overflow-hidden bg-zinc-900 border border-zinc-700/50 terminal-code-block">
      <div class="px-4 py-2.5 text-xs font-mono text-zinc-300 border-b border-zinc-700/50 bg-zinc-800/80">Terminal</div>
      <pre class="p-4 overflow-x-auto"><code class="font-mono text-sm text-white">pip install kubed
kubed-setup</code></pre>
    </div>
    <div class="rounded-xl overflow-hidden bg-zinc-900 border border-zinc-700/50 terminal-code-block">
      <div class="px-4 py-2.5 text-xs font-mono text-zinc-300 border-b border-zinc-700/50 bg-zinc-800/80">Non-interactive</div>
      <pre class="p-4 overflow-x-auto"><code class="font-mono text-sm text-white">pip install kubed
kubed-setup --force-yes</code></pre>
    </div>
  </div>
  <div class="mt-6 flex items-start gap-3 p-4 rounded-lg bg-zinc-100 border border-zinc-200 text-zinc-800">
    <span class="text-zinc-500 text-lg shrink-0" aria-hidden="true">ℹ</span>
    <p class="text-sm font-medium">Restart your terminal after installation for changes to take effect.</p>
  </div>
  <a href="/installation" class="inline-flex items-center gap-2 mt-6 px-4 py-2.5 bg-accent hover:bg-accent-hover text-white text-sm font-medium rounded-lg transition-colors">
    <i class="fas fa-book" aria-hidden="true"></i>
    Full installation guide
  </a>
</section>

<section class="text-center py-8 border-t border-zinc-200">
  <p class="text-zinc-600 mb-4">Found an error or have a suggestion?</p>
  <a href="https://github.com/dalefrieswthat/kubed/issues" class="inline-flex items-center gap-2 px-4 py-2 bg-zinc-800 hover:bg-zinc-900 text-white text-sm font-medium rounded-lg transition-colors">
    <i class="fas fa-bug" aria-hidden="true"></i>
    Report an issue
  </a>
</section>
