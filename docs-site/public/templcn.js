(function () {
  var searchIndex = [];
  var searchLoaded = false;
  var searchLoadPromise = null;
  var searchState = {
    open: false,
    activeIndex: 0,
    results: [],
  };

  var themePresets = {
    neutral: {
      label: 'Neutral',
      light: {
        radius: '0.625rem',
        background: 'oklch(1 0 0)',
        foreground: 'oklch(0.145 0 0)',
        card: 'oklch(1 0 0)',
        'card-foreground': 'oklch(0.145 0 0)',
        popover: 'oklch(1 0 0)',
        'popover-foreground': 'oklch(0.145 0 0)',
        primary: 'oklch(0.205 0 0)',
        'primary-foreground': 'oklch(0.985 0 0)',
        secondary: 'oklch(0.97 0 0)',
        'secondary-foreground': 'oklch(0.205 0 0)',
        muted: 'oklch(0.97 0 0)',
        'muted-foreground': 'oklch(0.556 0 0)',
        accent: 'oklch(0.97 0 0)',
        'accent-foreground': 'oklch(0.205 0 0)',
        destructive: 'oklch(0.577 0.245 27.325)',
        'destructive-foreground': 'oklch(0.985 0 0)',
        border: 'oklch(0.922 0 0)',
        input: 'oklch(0.922 0 0)',
        ring: 'oklch(0.708 0 0)',
        'chart-1': 'oklch(0.646 0.222 41.116)',
        'chart-2': 'oklch(0.6 0.118 184.704)',
        'chart-3': 'oklch(0.398 0.07 227.392)',
        'chart-4': 'oklch(0.828 0.189 84.429)',
        'chart-5': 'oklch(0.769 0.188 70.08)',
        sidebar: 'oklch(0.985 0 0)',
        'sidebar-foreground': 'oklch(0.145 0 0)',
        'sidebar-primary': 'oklch(0.205 0 0)',
        'sidebar-primary-foreground': 'oklch(0.985 0 0)',
        'sidebar-accent': 'oklch(0.97 0 0)',
        'sidebar-accent-foreground': 'oklch(0.205 0 0)',
        'sidebar-border': 'oklch(0.922 0 0)',
        'sidebar-ring': 'oklch(0.708 0 0)',
      },
      dark: {
        background: 'oklch(0.145 0 0)',
        foreground: 'oklch(0.985 0 0)',
        card: 'oklch(0.205 0 0)',
        'card-foreground': 'oklch(0.985 0 0)',
        popover: 'oklch(0.205 0 0)',
        'popover-foreground': 'oklch(0.985 0 0)',
        primary: 'oklch(0.922 0 0)',
        'primary-foreground': 'oklch(0.205 0 0)',
        secondary: 'oklch(0.269 0 0)',
        'secondary-foreground': 'oklch(0.985 0 0)',
        muted: 'oklch(0.269 0 0)',
        'muted-foreground': 'oklch(0.708 0 0)',
        accent: 'oklch(0.269 0 0)',
        'accent-foreground': 'oklch(0.985 0 0)',
        destructive: 'oklch(0.704 0.191 22.216)',
        'destructive-foreground': 'oklch(0.985 0 0)',
        border: 'oklch(1 0 0 / 10%)',
        input: 'oklch(1 0 0 / 15%)',
        ring: 'oklch(0.556 0 0)',
        'chart-1': 'oklch(0.488 0.243 264.376)',
        'chart-2': 'oklch(0.696 0.17 162.48)',
        'chart-3': 'oklch(0.769 0.188 70.08)',
        'chart-4': 'oklch(0.627 0.265 303.9)',
        'chart-5': 'oklch(0.645 0.246 16.439)',
        sidebar: 'oklch(0.205 0 0)',
        'sidebar-foreground': 'oklch(0.985 0 0)',
        'sidebar-primary': 'oklch(0.488 0.243 264.376)',
        'sidebar-primary-foreground': 'oklch(0.985 0 0)',
        'sidebar-accent': 'oklch(0.269 0 0)',
        'sidebar-accent-foreground': 'oklch(0.985 0 0)',
        'sidebar-border': 'oklch(1 0 0 / 10%)',
        'sidebar-ring': 'oklch(0.556 0 0)',
      },
    },
    zinc: {
      label: 'Zinc',
      extends: 'neutral',
      light: {
        foreground: 'oklch(0.141 0.005 285.823)',
        primary: 'oklch(0.21 0.006 285.885)',
        secondary: 'oklch(0.967 0.001 286.375)',
        muted: 'oklch(0.967 0.001 286.375)',
        accent: 'oklch(0.967 0.001 286.375)',
        border: 'oklch(0.92 0.004 286.32)',
        input: 'oklch(0.92 0.004 286.32)',
        ring: 'oklch(0.705 0.015 286.067)',
        sidebar: 'oklch(0.985 0 0)',
      },
      dark: {
        background: 'oklch(0.141 0.005 285.823)',
        card: 'oklch(0.21 0.006 285.885)',
        popover: 'oklch(0.21 0.006 285.885)',
        secondary: 'oklch(0.274 0.006 286.033)',
        muted: 'oklch(0.274 0.006 286.033)',
        accent: 'oklch(0.274 0.006 286.033)',
        ring: 'oklch(0.552 0.016 285.938)',
        sidebar: 'oklch(0.21 0.006 285.885)',
      },
    },
    stone: {
      label: 'Stone',
      extends: 'neutral',
      light: {
        foreground: 'oklch(0.147 0.004 49.25)',
        primary: 'oklch(0.216 0.006 56.043)',
        secondary: 'oklch(0.97 0.001 106.424)',
        muted: 'oklch(0.97 0.001 106.424)',
        accent: 'oklch(0.97 0.001 106.424)',
        border: 'oklch(0.923 0.003 48.717)',
        input: 'oklch(0.923 0.003 48.717)',
        ring: 'oklch(0.709 0.01 56.259)',
        'chart-4': 'oklch(0.828 0.189 84.429)',
      },
      dark: {
        background: 'oklch(0.147 0.004 49.25)',
        card: 'oklch(0.216 0.006 56.043)',
        popover: 'oklch(0.216 0.006 56.043)',
        secondary: 'oklch(0.268 0.007 34.298)',
        muted: 'oklch(0.268 0.007 34.298)',
        accent: 'oklch(0.268 0.007 34.298)',
        ring: 'oklch(0.553 0.013 58.071)',
      },
    },
    blue: {
      label: 'Blue',
      extends: 'neutral',
      light: {
        primary: 'oklch(0.546 0.245 262.881)',
        'primary-foreground': 'oklch(0.985 0 0)',
        ring: 'oklch(0.623 0.214 259.815)',
        accent: 'oklch(0.97 0.014 254.604)',
        'accent-foreground': 'oklch(0.28 0.091 267.935)',
        'chart-1': 'oklch(0.623 0.214 259.815)',
        'chart-2': 'oklch(0.707 0.165 254.624)',
        'sidebar-primary': 'oklch(0.546 0.245 262.881)',
      },
      dark: {
        primary: 'oklch(0.707 0.165 254.624)',
        'primary-foreground': 'oklch(0.21 0.006 285.885)',
        ring: 'oklch(0.488 0.243 264.376)',
        accent: 'oklch(0.379 0.146 265.522)',
        'accent-foreground': 'oklch(0.985 0 0)',
        'chart-1': 'oklch(0.707 0.165 254.624)',
        'chart-2': 'oklch(0.623 0.214 259.815)',
        'sidebar-primary': 'oklch(0.707 0.165 254.624)',
      },
    },
  };

  function mergedPreset(slug) {
    var preset = themePresets[slug] || themePresets.neutral;
    if (!preset.extends) return preset;
    var base = mergedPreset(preset.extends);
    return {
      label: preset.label,
      light: Object.assign({}, base.light, preset.light || {}),
      dark: Object.assign({}, base.dark, preset.dark || {}),
    };
  }

  function declarations(vars) {
    return Object.keys(vars).map(function (key) {
      return '  --' + key + ': ' + vars[key] + ';';
    }).join('\n');
  }

  function themeTokenCSS(slug, scoped) {
    var preset = mergedPreset(slug);
    var rootSelector = scoped ? 'html[data-theme-preset="' + slug + '"]' : ':root';
    var darkSelector = scoped ? 'html.dark[data-theme-preset="' + slug + '"]' : '.dark';
    return rootSelector + ' {\n' + declarations(preset.light) + '\n}\n\n' + darkSelector + ' {\n' + declarations(preset.dark) + '\n}';
  }

  function globalsCSS(slug) {
    return '@import "tailwindcss";\n@import "tw-animate-css";\n\n@source "./**/*.go";\n@source "./**/*.templ";\n\n@custom-variant dark (&:is(.dark *));\n\n@theme inline {\n  --color-background: var(--background);\n  --color-foreground: var(--foreground);\n  --color-card: var(--card);\n  --color-card-foreground: var(--card-foreground);\n  --color-popover: var(--popover);\n  --color-popover-foreground: var(--popover-foreground);\n  --color-primary: var(--primary);\n  --color-primary-foreground: var(--primary-foreground);\n  --color-secondary: var(--secondary);\n  --color-secondary-foreground: var(--secondary-foreground);\n  --color-muted: var(--muted);\n  --color-muted-foreground: var(--muted-foreground);\n  --color-accent: var(--accent);\n  --color-accent-foreground: var(--accent-foreground);\n  --color-destructive: var(--destructive);\n  --color-destructive-foreground: var(--destructive-foreground);\n  --color-border: var(--border);\n  --color-input: var(--input);\n  --color-ring: var(--ring);\n  --color-chart-1: var(--chart-1);\n  --color-chart-2: var(--chart-2);\n  --color-chart-3: var(--chart-3);\n  --color-chart-4: var(--chart-4);\n  --color-chart-5: var(--chart-5);\n  --color-sidebar: var(--sidebar);\n  --color-sidebar-foreground: var(--sidebar-foreground);\n  --color-sidebar-primary: var(--sidebar-primary);\n  --color-sidebar-primary-foreground: var(--sidebar-primary-foreground);\n  --color-sidebar-accent: var(--sidebar-accent);\n  --color-sidebar-accent-foreground: var(--sidebar-accent-foreground);\n  --color-sidebar-border: var(--sidebar-border);\n  --color-sidebar-ring: var(--sidebar-ring);\n  --radius-sm: calc(var(--radius) * 0.6);\n  --radius-md: calc(var(--radius) * 0.8);\n  --radius-lg: var(--radius);\n  --radius-xl: calc(var(--radius) * 1.4);\n  --radius-2xl: calc(var(--radius) * 1.8);\n  --radius-3xl: calc(var(--radius) * 2.2);\n  --radius-4xl: calc(var(--radius) * 2.6);\n}\n\n' + themeTokenCSS(slug, false) + '\n\n@layer base {\n  * {\n    @apply border-border outline-ring/50;\n  }\n\n  body {\n    @apply bg-background text-foreground;\n  }\n}\n';
  }

  function setThemeMode(mode) {
    var root = document.documentElement;
    var resolved = mode;
    if (mode === 'system') {
      resolved = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }
    root.classList.toggle('dark', resolved === 'dark');
    root.dataset.themeMode = mode;
    try {
      localStorage.setItem('theme', mode);
    } catch (_) {}
    updateThemeControls();
  }

  function setTheme(theme) {
    setThemeMode(theme);
  }

  function setThemePreset(slug) {
    if (!themePresets[slug]) slug = 'neutral';
    var style = document.getElementById('theme-preset-style');
    if (!style) {
      style = document.createElement('style');
      style.id = 'theme-preset-style';
      document.head.appendChild(style);
    }
    style.textContent = themeTokenCSS(slug, true);
    document.documentElement.dataset.themePreset = slug;
    try {
      localStorage.setItem('theme-preset', slug);
    } catch (_) {}
    updateThemeControls();
  }

  function initTheme() {
    var savedMode = 'system';
    var savedPreset = 'neutral';
    try {
      var saved = localStorage.getItem('theme');
      if (saved === 'dark' || saved === 'light' || saved === 'system') {
        savedMode = saved;
      }
      savedPreset = localStorage.getItem('theme-preset') || 'neutral';
    } catch (_) {}

    setThemePreset(savedPreset);
    setThemeMode(savedMode);
    if (window.matchMedia) {
      window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function () {
        if (document.documentElement.dataset.themeMode === 'system') setThemeMode('system');
      });
    }
  }

  function updateThemeControls() {
    var root = document.documentElement;
    var mode = root.dataset.themeMode || 'system';
    var preset = root.dataset.themePreset || 'neutral';
    document.querySelectorAll('[data-theme-mode]').forEach(function (button) {
      var active = button.getAttribute('data-theme-mode') === mode;
      button.setAttribute('data-state', active ? 'active' : 'inactive');
      button.classList.toggle('bg-accent', active);
      button.classList.toggle('text-foreground', active);
    });
    document.querySelectorAll('[data-theme-preset]').forEach(function (button) {
      var active = button.getAttribute('data-theme-preset') === preset;
      button.setAttribute('data-state', active ? 'active' : 'inactive');
      button.classList.toggle('ring-2', active);
      button.classList.toggle('ring-ring', active);
      button.classList.toggle('bg-accent', active);
    });
  }

  function initThemeToggles() {
    document.querySelectorAll('[data-theme-toggle]').forEach(function (button) {
      button.addEventListener('click', function () {
        var root = document.documentElement;
        setThemeMode(root.classList.contains('dark') ? 'light' : 'dark');
      });
    });

    document.querySelectorAll('[data-theme-menu]').forEach(function (menu) {
      var trigger = menu.querySelector('[data-theme-menu-trigger]');
      var content = menu.querySelector('[data-theme-menu-content]');
      if (!trigger || !content) return;
      trigger.addEventListener('click', function (event) {
        event.stopPropagation();
        var isHidden = content.classList.toggle('hidden');
        trigger.setAttribute('aria-expanded', isHidden ? 'false' : 'true');
      });
      content.addEventListener('click', function (event) {
        event.stopPropagation();
      });
    });

    document.addEventListener('click', function () {
      document.querySelectorAll('[data-theme-menu-content]').forEach(function (content) {
        content.classList.add('hidden');
      });
      document.querySelectorAll('[data-theme-menu-trigger]').forEach(function (trigger) {
        trigger.setAttribute('aria-expanded', 'false');
      });
    });

    document.querySelectorAll('[data-theme-mode]').forEach(function (button) {
      button.addEventListener('click', function () {
        setThemeMode(button.getAttribute('data-theme-mode') || 'system');
      });
    });

    document.querySelectorAll('[data-theme-preset]').forEach(function (button) {
      button.addEventListener('click', function () {
        setThemePreset(button.getAttribute('data-theme-preset') || 'neutral');
      });
    });

    document.querySelectorAll('[data-theme-copy]').forEach(function (button) {
      button.addEventListener('click', async function () {
        var slug = document.documentElement.dataset.themePreset || 'neutral';
        try {
          await navigator.clipboard.writeText(globalsCSS(slug));
          button.textContent = 'Copied';
          window.setTimeout(function () { button.textContent = button.hasAttribute('data-theme-copy') && button.closest('[data-theme-menu-content]') ? 'Copy CSS' : 'Copy current CSS'; }, 1200);
        } catch (_) {
          button.textContent = 'Copy failed';
        }
      });
    });

    document.querySelectorAll('[data-theme-download]').forEach(function (button) {
      button.addEventListener('click', function () {
        var slug = document.documentElement.dataset.themePreset || 'neutral';
        var blob = new Blob([globalsCSS(slug)], { type: 'text/css;charset=utf-8' });
        var url = URL.createObjectURL(blob);
        var link = document.createElement('a');
        link.href = url;
        link.download = 'globals.css';
        document.body.appendChild(link);
        link.click();
        link.remove();
        URL.revokeObjectURL(url);
      });
    });

    updateThemeControls();
  }



  function loadSearchIndex() {
    if (searchLoaded) {
      return Promise.resolve(searchIndex);
    }
    if (searchLoadPromise) {
      return searchLoadPromise;
    }

    searchLoadPromise = fetch('/search-index.json')
      .then(function (response) {
        if (!response.ok) {
          throw new Error('Failed to load search index');
        }
        return response.json();
      })
      .then(function (data) {
        searchIndex = Array.isArray(data) ? data : [];
        searchLoaded = true;
        return searchIndex;
      })
      .catch(function (error) {
        console.error('Error loading search index:', error);
        searchIndex = [];
        searchLoaded = false;
        searchLoadPromise = null;
        return searchIndex;
      });

    return searchLoadPromise;
  }

  function searchTypeLabel(type) {
    if (type === 'component') return 'Component';
    if (type === 'block') return 'Block';
    return 'Doc';
  }

  function searchTypeClass(type) {
    if (type === 'component') return 'bg-primary/10 text-primary';
    if (type === 'block') return 'bg-secondary/15 text-secondary-foreground';
    return 'bg-accent/10 text-accent-foreground';
  }

  function renderSearchResults(query) {
    var resultsContainer = document.querySelector('[data-search-results="true"]');
    if (!resultsContainer) return;

    var trimmed = (query || '').trim().toLowerCase();
    if (!trimmed) {
      searchState.results = [];
      searchState.activeIndex = 0;
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">Type to search...</p>';
      return;
    }

    var matches = searchIndex
      .filter(function (entry) {
        var title = String(entry.title || '').toLowerCase();
        var url = String(entry.url || '').toLowerCase();
        var type = String(entry.type || '').toLowerCase();
        return title.indexOf(trimmed) !== -1 || url.indexOf(trimmed) !== -1 || type.indexOf(trimmed) !== -1;
      })
      .slice(0, 12);

    searchState.results = matches;
    searchState.activeIndex = matches.length > 0 ? 0 : -1;

    if (matches.length === 0) {
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">No results found</p>';
      return;
    }

    var html = '<ul class="flex flex-col gap-1">';
    matches.forEach(function (entry, index) {
      var selected = index === searchState.activeIndex;
      html +=
        '<li>' +
        '<a data-search-result="true" data-search-index="' +
        index +
        '" href="' +
        entry.url +
        '" class="flex items-center justify-between gap-4 rounded-md px-3 py-2 text-sm ' +
        (selected ? 'bg-accent text-accent-foreground' : 'hover:bg-accent hover:text-accent-foreground') +
        '">' +
        '<span>' +
        entry.title +
        '</span>' +
        '<span class="rounded px-1.5 py-0.5 text-xs ' +
        searchTypeClass(entry.type) +
        '">' +
        searchTypeLabel(entry.type) +
        '</span>' +
        '</a>' +
        '</li>';
    });
    html += '</ul>';
    resultsContainer.innerHTML = html;
  }

  function setActiveSearchResult(index) {
    var resultsContainer = document.querySelector('[data-search-results="true"]');
    if (!resultsContainer || !searchState.results.length) return;

    var max = searchState.results.length - 1;
    searchState.activeIndex = Math.max(0, Math.min(index, max));

    resultsContainer.querySelectorAll('[data-search-result="true"]').forEach(function (item) {
      var current = Number(item.getAttribute('data-search-index'));
      var active = current === searchState.activeIndex;
      item.classList.toggle('bg-accent', active);
      item.classList.toggle('text-accent-foreground', active);
      item.scrollIntoView({ block: 'nearest' });
    });
  }

  function openSearchDialog() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    if (!dialog || !input) return;

    dialog.classList.remove('hidden');
    searchState.open = true;
    window.requestAnimationFrame(function () {
      input.focus();
      input.select();
    });
  }

  function closeSearchDialog() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    var resultsContainer = document.querySelector('[data-search-results="true"]');

    if (dialog) {
      dialog.classList.add('hidden');
    }
    if (input) {
      input.value = '';
    }
    if (resultsContainer) {
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">Type to search...</p>';
    }
    searchState.open = false;
    searchState.activeIndex = 0;
    searchState.results = [];
  }

  function navigateSearchResult(delta) {
    if (!searchState.results.length) return;
    var next = searchState.activeIndex + delta;
    if (next < 0) next = searchState.results.length - 1;
    if (next >= searchState.results.length) next = 0;
    setActiveSearchResult(next);
  }

  function initSearch() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var backdrop = document.querySelector('[data-search-backdrop="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    var searchResults = document.querySelector('[data-search-results="true"]');

    if (!dialog || !input) return;

    document.addEventListener('keydown', function (event) {
      if ((event.metaKey || event.ctrlKey) && event.key.toLowerCase() === 'k') {
        event.preventDefault();
        if (dialog.classList.contains('hidden')) {
          openSearchDialog();
          loadSearchIndex().then(function () {
            renderSearchResults(input.value);
          });
        } else {
          closeSearchDialog();
        }
        return;
      }

      if (!searchState.open) return;

      if (event.key === 'Escape') {
        event.preventDefault();
        closeSearchDialog();
        return;
      }

      if (event.key === 'ArrowDown') {
        event.preventDefault();
        navigateSearchResult(1);
        return;
      }

      if (event.key === 'ArrowUp') {
        event.preventDefault();
        navigateSearchResult(-1);
        return;
      }

      if (event.key === 'Enter' && searchState.results.length > 0) {
        event.preventDefault();
        var selected = searchState.results[searchState.activeIndex];
        if (selected && selected.url) {
          window.location.assign(selected.url);
        }
      }
    });

    if (backdrop) {
      backdrop.addEventListener('click', closeSearchDialog);
    }

    document.querySelectorAll('[data-search-trigger="true"]').forEach(function (button) {
      button.addEventListener('click', function () {
        if (dialog.classList.contains('hidden')) {
          openSearchDialog();
          loadSearchIndex().then(function () {
            renderSearchResults(input.value);
          });
        } else {
          closeSearchDialog();
        }
      });
    });

    input.addEventListener('input', function (event) {
      var query = event.target.value;
      loadSearchIndex().then(function () {
        renderSearchResults(query);
      });
    });

    if (searchResults) {
      searchResults.addEventListener('mousemove', function (event) {
        var link = event.target.closest('[data-search-result="true"]');
        if (!link) return;
        var index = Number(link.getAttribute('data-search-index'));
        if (!Number.isNaN(index)) {
          searchState.activeIndex = index;
          setActiveSearchResult(index);
        }
      });

      searchResults.addEventListener('click', function (event) {
        var link = event.target.closest('a');
        if (link) {
          closeSearchDialog();
        }
      });
    }
  }

  function initMobileMenu() {
    var toggle = document.querySelector('[data-mobile-menu-toggle]');
    var drawer = document.querySelector('[data-mobile-drawer]');
    var close = document.querySelector('[data-mobile-drawer-close]');
    if (!toggle || !drawer) return;

    function openDrawer() {
      drawer.classList.remove('hidden');
      drawer.classList.add('flex');
      window.requestAnimationFrame(function () {
        drawer.classList.remove('translate-x-full');
        drawer.classList.add('translate-x-0');
      });
    }

    function closeDrawer() {
      drawer.classList.add('translate-x-full');
      drawer.classList.remove('translate-x-0');
      window.setTimeout(function () {
        drawer.classList.add('hidden');
        drawer.classList.remove('flex');
      }, 200);
    }

    toggle.addEventListener('click', function () {
      if (drawer.classList.contains('hidden')) openDrawer();
      else closeDrawer();
    });

    if (close) {
      close.addEventListener('click', closeDrawer);
    }

    drawer.addEventListener('click', function (event) {
      if (event.target === drawer) {
        closeDrawer();
      }
    });

    drawer.querySelectorAll('a').forEach(function (link) {
      link.addEventListener('click', closeDrawer);
    });
  }

  function initCopyButtons() {
    document.querySelectorAll('pre').forEach(function (pre) {
      if (pre.closest('[data-preview-card="true"]')) return;
      if (pre.dataset.copyReady === 'true') return;
      pre.dataset.copyReady = 'true';
      pre.classList.add('relative');

      var button = document.createElement('button');
      button.type = 'button';
      button.textContent = 'Copy';
      button.className =
        'absolute right-2 top-2 rounded-md border border-border/70 bg-background px-2 py-1 text-xs font-medium text-muted-foreground transition-colors hover:text-foreground';

      button.addEventListener('click', async function () {
        try {
          await navigator.clipboard.writeText(pre.querySelector('code') ? pre.querySelector('code').textContent : pre.textContent);
          button.textContent = 'Copied';
          window.setTimeout(function () {
            button.textContent = 'Copy';
          }, 1200);
        } catch (_) {
          button.textContent = 'Copy failed';
        }
      });

      pre.appendChild(button);
    });
  }

  function initComponentToc() {
    document.querySelectorAll('[data-toc-root="true"]').forEach(function (root) {
      var sections = Array.prototype.slice.call(root.querySelectorAll('#installation, #usage, #composition, #examples, #api-reference, #view-source'));
      var links = Array.prototype.slice.call(root.querySelectorAll('[data-toc-link]'));
      if (!sections.length || !links.length) return;

      function updateActive() {
        var activeId = sections[0].id;
        var offset = window.innerHeight * 0.3;
        sections.forEach(function (section) {
          var rect = section.getBoundingClientRect();
          if (rect.top <= offset) activeId = section.id;
        });
        links.forEach(function (link) {
          var isActive = link.getAttribute('data-toc-link') === activeId;
          link.classList.toggle('bg-accent', isActive);
          link.classList.toggle('text-foreground', isActive);
          link.classList.toggle('text-muted-foreground', !isActive);
        });
      }

      updateActive();
      window.addEventListener('scroll', updateActive);
    });
  }

  function registerTemplHighlighting() {
    if (!window.hljs || typeof window.hljs.registerLanguage !== 'function' || window.__templHighlightRegistered) return;
    window.__templHighlightRegistered = true;
    window.hljs.registerLanguage('templ', function (hljs) {
      return {
        name: 'templ',
        aliases: ['gotempl'],
        contains: [
          {
            className: 'title function_',
            begin: /@[A-Za-z_][A-Za-z0-9_.]*(?=\()/,
          },
          {
            className: 'type',
            begin: /\b[A-Za-z_][A-Za-z0-9_]*(Props|Variant|Size)?\b(?=\{)/,
          },
          {
            className: 'attr',
            begin: /\b[A-Za-z_][A-Za-z0-9_]*(?=:)/,
          },
          {
            className: 'tag',
            begin: /<\/?[A-Za-z][^>]*>/,
          },
          hljs.QUOTE_STRING_MODE,
          hljs.C_NUMBER_MODE,
          hljs.C_LINE_COMMENT_MODE,
          hljs.C_BLOCK_COMMENT_MODE,
          {
            className: 'literal',
            begin: /\b(true|false|nil)\b/,
          },
        ],
      };
    });
  }

  function initHighlighting() {
    function runHighlight() {
      if (window.hljs && typeof window.hljs.highlightAll === 'function') {
        registerTemplHighlighting();
        // Only highlight elements that are visible (not hidden by a parent).
        // Hidden code panes in PreviewCard are highlighted on-demand when the
        // Code tab opens, so we skip them here to avoid double-processing.
        document.querySelectorAll('pre code:not([data-highlighted]):not([data-raw-code])').forEach(function(el) {
          if (!el.closest('[data-preview-panel="code"]')) {
            window.hljs.highlightElement(el);
          }
        });
        return true;
      }
      return false;
    }
    // hljs is loaded synchronously in <head> so it should be available immediately.
    // Retry with rAF as a safety net for slow CDN connections.
    if (!runHighlight()) {
      var attempts = 0;
      var timer = setInterval(function() {
        if (runHighlight() || ++attempts > 20) clearInterval(timer);
      }, 100);
    }
  }

  document.addEventListener('DOMContentLoaded', function () {
    initTheme();
    initThemeToggles();
    initMobileMenu();
    initSearch();
    initCopyButtons();
    initComponentToc();
    initHighlighting();
  });
})();
