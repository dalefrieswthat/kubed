---
layout: default
title: Kubed Documentation
nav_order: 1
description: "Documentation and cheat sheets for Docker, Kubernetes, Terraform, and Helm"
permalink: /
---

<div class="hero">
  <h1 class="animate-title">Kubed <span>Documentation</span></h1>
  <p class="lead">Your comprehensive guide to cloud-native development tools</p>
  <div class="hero-pattern"></div>
</div>

<div class="features">
  <div class="feature-card" data-tool="docker">
    <div class="feature-icon">
      <i class="fab fa-docker"></i>
    </div>
    <h2>Docker</h2>
    <p>Container management and orchestration</p>
    <a href="/docker" class="button">View Commands</a>
  </div>

  <div class="feature-card" data-tool="kubernetes">
    <div class="feature-icon">
      <i class="fas fa-dharmachakra"></i>
    </div>
    <h2>Kubernetes</h2>
    <p>Container orchestration and management</p>
    <a href="/kubernetes" class="button">View Commands</a>
  </div>

  <div class="feature-card" data-tool="terraform">
    <div class="feature-icon">
      <i class="fas fa-cube"></i>
    </div>
    <h2>Terraform</h2>
    <p>Infrastructure as Code</p>
    <a href="/terraform" class="button">View Commands</a>
  </div>

  <div class="feature-card" data-tool="helm">
    <div class="feature-icon">
      <i class="fas fa-chart-pie"></i>
    </div>
    <h2>Helm</h2>
    <p>Kubernetes package management</p>
    <a href="/helm" class="button">View Commands</a>
  </div>
</div>

<div class="cta">
  <h2>Get Started with Kubed</h2>
  <p>Enhance your terminal experience with powerful autocompletion and useful aliases</p>
  <div class="code-block">
    <pre><code class="language-bash">pip install kubed
kubed-setup</code></pre>
  </div>
  <div class="code-block">
    <h4>For non-interactive installation:</h4>
    <pre><code class="language-bash">pip install kubed
kubed-setup --force-yes</code></pre>
  </div>
  <div class="warning">
    <span class="warning-icon">⚠️</span> You MUST restart your terminal after installation for changes to take effect!
  </div>
  <div class="learn-more">
    <a href="/installation" class="button">
      <i class="fas fa-book"></i> View Complete Installation Guide
    </a>
  </div>
</div>

<div class="footer-cta">
  <p>Found an error or have a suggestion?</p>
  <a href="https://github.com/dalefrieswthat/kubed/issues" class="button">
    <i class="fas fa-bug"></i> Report an Issue
  </a>
</div> 