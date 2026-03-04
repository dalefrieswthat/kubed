---
layout: default
title: Home
nav_order: 1
description: "Documentation and cheat sheets for Docker, Kubernetes, Terraform, and Helm"
permalink: /
---

<section class="mb-16">
  <h1 class="text-4xl md:text-5xl font-semibold text-surface-900 tracking-tight mb-4">
    Kubed
  </h1>
  <p class="text-xl text-slate-700 max-w-2xl leading-relaxed">
    CLI productivity tool with autocompletion for Docker, Kubernetes, Terraform, and Helm. Your reference for cloud-native development.
  </p>
</section>

<section class="mb-16">
  <h2 class="text-sm font-semibold text-slate-600 uppercase tracking-wider mb-6">Tools</h2>
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
    <a href="/docker" class="group block bg-white rounded-2xl border border-surface-200 p-6 shadow-card hover:shadow-soft hover:border-primary-500/30 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#2496ED]/10 flex items-center justify-center mb-4 group-hover:bg-[#2496ED]/20 transition-colors">
        <i class="fab fa-docker text-[#2496ED] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-surface-900 mb-1">Docker</h3>
      <p class="text-sm text-slate-600">Container management and orchestration</p>
    </a>
    <a href="/kubernetes" class="group block bg-white rounded-2xl border border-surface-200 p-6 shadow-card hover:shadow-soft hover:border-primary-500/30 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#326CE5]/10 flex items-center justify-center mb-4 group-hover:bg-[#326CE5]/20 transition-colors">
        <i class="fas fa-dharmachakra text-[#326CE5] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-surface-900 mb-1">Kubernetes</h3>
      <p class="text-sm text-slate-600">Container orchestration and management</p>
    </a>
    <a href="/terraform" class="group block bg-white rounded-2xl border border-surface-200 p-6 shadow-card hover:shadow-soft hover:border-primary-500/30 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#7B42BC]/10 flex items-center justify-center mb-4 group-hover:bg-[#7B42BC]/20 transition-colors">
        <i class="fas fa-cube text-[#7B42BC] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-surface-900 mb-1">Terraform</h3>
      <p class="text-sm text-slate-600">Infrastructure as Code</p>
    </a>
    <a href="/helm" class="group block bg-white rounded-2xl border border-surface-200 p-6 shadow-card hover:shadow-soft hover:border-primary-500/30 transition-all duration-200">
      <div class="w-10 h-10 rounded-lg bg-[#0F1689]/10 flex items-center justify-center mb-4 group-hover:bg-[#0F1689]/20 transition-colors">
        <i class="fas fa-chart-pie text-[#0F1689] text-xl" aria-hidden="true"></i>
      </div>
      <h3 class="text-lg font-semibold text-surface-900 mb-1">Helm</h3>
      <p class="text-sm text-slate-500">Kubernetes package management</p>
    </a>
  </div>
</section>

<section class="bg-white rounded-2xl border border-surface-200 p-8 md:p-10 shadow-card mb-12">
  <h2 class="text-2xl font-semibold text-surface-900 mb-2">Get started</h2>
  <p class="text-slate-700 mb-6">Install Kubed and run setup for your shell. Restart your terminal after installation for changes to take effect.</p>
  <div class="space-y-4">
    <div class="rounded-xl bg-surface-900 overflow-hidden">
      <div class="px-4 py-2 text-xs font-mono text-slate-300 border-b border-white/10">Terminal</div>
      <pre class="p-4 overflow-x-auto"><code class="font-mono text-sm text-slate-200">pip install kubed
kubed-setup</code></pre>
    </div>
    <div class="rounded-xl bg-surface-900 overflow-hidden">
      <div class="px-4 py-2 text-xs font-mono text-slate-300 border-b border-white/10">Non-interactive</div>
      <pre class="p-4 overflow-x-auto"><code class="font-mono text-sm text-slate-200">pip install kubed
kubed-setup --force-yes</code></pre>
    </div>
  </div>
  <div class="mt-6 flex items-start gap-3 p-4 rounded-lg bg-amber-50 border border-amber-200/60">
    <span class="text-amber-600 text-lg" aria-hidden="true">⚠️</span>
    <p class="text-sm text-amber-800 font-medium">Restart your terminal after installation for changes to take effect.</p>
  </div>
  <a href="/installation" class="inline-flex items-center gap-2 mt-6 px-4 py-2.5 bg-primary-600 hover:bg-primary-700 text-white text-sm font-medium rounded-lg transition-colors">
    <i class="fas fa-book" aria-hidden="true"></i>
    Full installation guide
  </a>
</section>

<section class="text-center py-8 border-t border-surface-200">
  <p class="text-slate-600 mb-4">Found an error or have a suggestion?</p>
  <a href="https://github.com/dalefrieswthat/kubed/issues" class="inline-flex items-center gap-2 px-4 py-2 bg-surface-800 hover:bg-surface-900 text-white text-sm font-medium rounded-lg transition-colors">
    <i class="fas fa-bug" aria-hidden="true"></i>
    Report an issue
  </a>
</section>
