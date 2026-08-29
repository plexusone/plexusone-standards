// Example file with design system violations
// Used for demonstration purposes

import React from 'react';

// VIOLATION: Hardcoded colors (should use tokens)
const styles = {
  container: {
    background: '#0a0e1a',  // Should be: var(--plexus-dark)
    padding: '24px',        // Should be: var(--plexus-space-6)
  },
  heading: {
    color: '#f1f5f9',       // Should be: var(--plexus-text)
    fontFamily: 'Inter, sans-serif',  // Should be: var(--plexus-font-sans)
    marginBottom: '16px',   // Should be: var(--plexus-space-4)
  },
  card: {
    background: '#0f172a',  // Should be: var(--plexus-navy)
    borderRadius: '8px',    // Should be: var(--plexus-radius-lg)
    padding: '32px',        // Should be: var(--plexus-space-8)
  },
  link: {
    color: '#06b6d4',       // Should be: var(--plexus-cyan)
  },
  muted: {
    color: '#94a3b8',       // Should be: var(--plexus-text-muted)
  },
};

// VIOLATION: Multiple primary buttons in same view
function ActionButtons() {
  return (
    <div style={{ display: 'flex', gap: '8px' }}>
      <button
        style={{
          background: '#06b6d4',  // Primary color
          color: 'white',
          padding: '8px 16px',
          borderRadius: '4px',
        }}
      >
        Save Changes
      </button>
      <button
        style={{
          background: '#06b6d4',  // VIOLATION: Should be secondary
          color: 'white',
          padding: '8px 16px',
          borderRadius: '4px',
        }}
      >
        Cancel
      </button>
    </div>
  );
}

// VIOLATION: Missing focus and hover states
function NavLink({ href, children }: { href: string; children: React.ReactNode }) {
  return (
    <a
      href={href}
      style={{
        color: '#06b6d4',
        textDecoration: 'none',
        // Missing: hover state with cyan-light
        // Missing: focus-visible state with purple ring
      }}
    >
      {children}
    </a>
  );
}

// VIOLATION: Cyan used for large background area
function HeroSection() {
  return (
    <section
      style={{
        background: '#06b6d4',  // VIOLATION: Cyan is for accents, not backgrounds
        padding: '64px',
        color: 'white',
      }}
    >
      <h1>Welcome to PlexusOne</h1>
      <p>AI Agent Development Platform</p>
    </section>
  );
}

// VIOLATION: Arbitrary spacing not from scale
function CardGrid() {
  return (
    <div
      style={{
        display: 'grid',
        gap: '18px',  // VIOLATION: Not in spacing scale (should be 16px or 24px)
        padding: '22px',  // VIOLATION: Not in spacing scale
      }}
    >
      <div style={styles.card}>Card 1</div>
      <div style={styles.card}>Card 2</div>
      <div style={styles.card}>Card 3</div>
    </div>
  );
}

// COMPLIANT: Correct usage example
function CompliantExample() {
  return (
    <div
      style={{
        background: 'var(--plexus-dark)',
        padding: 'var(--plexus-space-6)',
        color: 'var(--plexus-text)',
        fontFamily: 'var(--plexus-font-sans)',
      }}
    >
      <h2 style={{ color: 'var(--plexus-cyan)' }}>
        Correct Token Usage
      </h2>
      <p style={{ color: 'var(--plexus-text-muted)' }}>
        This section uses design tokens correctly.
      </p>
    </div>
  );
}

export { ActionButtons, NavLink, HeroSection, CardGrid, CompliantExample };
