---
layout: default
title: Home
nav_order: 1
description: "Documentation and cheat sheets for Docker, Kubernetes, Terraform, and Helm"
permalink: /
---

<div class="hero bg-gradient-to-r from-blue-600 to-indigo-800 text-white py-16 px-4 rounded-lg shadow-xl mb-12">
  <h1 class="text-4xl md:text-5xl font-bold mb-4">Kubed <span class="text-blue-200">Documentation</span></h1>
  <p class="text-xl opacity-90 max-w-2xl">Your comprehensive guide to cloud-native development tools</p>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
  <div class="feature-card bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 border-t-4 border-blue-500">
    <div class="feature-icon text-blue-500 text-4xl mb-4">
      <i class="fab fa-docker"></i>
    </div>
    <h2 class="text-2xl font-bold mb-2">Docker</h2>
    <p class="text-gray-600 mb-4">Container management and orchestration</p>
    <a href="/docker" class="inline-block bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded transition-colors duration-300">View Commands</a>
  </div>

  <div class="feature-card bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 border-t-4 border-indigo-500">
    <div class="feature-icon text-indigo-500 text-4xl mb-4">
      <i class="fas fa-dharmachakra"></i>
    </div>
    <h2 class="text-2xl font-bold mb-2">Kubernetes</h2>
    <p class="text-gray-600 mb-4">Container orchestration and management</p>
    <a href="/kubernetes" class="inline-block bg-indigo-500 hover:bg-indigo-600 text-white py-2 px-4 rounded transition-colors duration-300">View Commands</a>
  </div>

  <div class="feature-card bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 border-t-4 border-purple-500">
    <div class="feature-icon text-purple-500 text-4xl mb-4">
      <i class="fas fa-cube"></i>
    </div>
    <h2 class="text-2xl font-bold mb-2">Terraform</h2>
    <p class="text-gray-600 mb-4">Infrastructure as Code automation</p>
    <a href="/terraform" class="inline-block bg-purple-500 hover:bg-purple-600 text-white py-2 px-4 rounded transition-colors duration-300">View Commands</a>
  </div>

  <div class="feature-card bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 border-t-4 border-teal-500">
    <div class="feature-icon text-teal-500 text-4xl mb-4">
      <i class="fas fa-chart-pie"></i>
    </div>
    <h2 class="text-2xl font-bold mb-2">Helm</h2>
    <p class="text-gray-600 mb-4">Kubernetes package management</p>
    <a href="/helm" class="inline-block bg-teal-500 hover:bg-teal-600 text-white py-2 px-4 rounded transition-colors duration-300">View Commands</a>
  </div>
</div>

<div class="cta bg-gray-100 rounded-lg p-8 mb-12">
  <h2 class="text-3xl font-bold mb-4">Get Started with Kubed</h2>
  <p class="text-xl text-gray-700 mb-6">Enhance your terminal experience with powerful autocompletion and useful aliases</p>
  <div class="code-block bg-gray-900 rounded-lg p-4 mb-4">
    <pre class="text-green-400"><code class="language-bash">pip install kubed
kubed-setup</code></pre>
  </div>
  <div class="code-block bg-gray-900 rounded-lg p-4 mb-4">
    <h4 class="text-white mb-2">For non-interactive installation:</h4>
    <pre class="text-green-400"><code class="language-bash">pip install kubed
kubed-setup --force-yes</code></pre>
  </div>
  <div class="warning bg-yellow-100 border-l-4 border-yellow-500 p-4 text-yellow-700 mb-6">
    <span class="warning-icon text-2xl">⚠️</span> You MUST restart your terminal after installation for changes to take effect!
  </div>
  <div class="learn-more">
    <a href="/installation" class="inline-block bg-blue-600 hover:bg-blue-700 text-white py-3 px-6 rounded-lg shadow-md hover:shadow-lg transition duration-300 font-medium">
      <i class="fas fa-book mr-2"></i> View Complete Installation Guide
    </a>
  </div>
</div>

<div class="footer-cta text-center py-8">
  <p class="text-gray-600 mb-4">Found an error or have a suggestion?</p>
  <a href="https://github.com/dalefrieswthat/kubed/issues" class="inline-block bg-gray-800 hover:bg-gray-900 text-white py-2 px-4 rounded-lg transition duration-300">
    <i class="fas fa-bug mr-2"></i> Report an Issue
  </a>
</div> 