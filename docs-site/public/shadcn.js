(function () {
  var searchIndex = [];
  var searchLoaded = false;
  var searchLoadPromise = null;
  var searchState = {
    open: false,
    activeIndex: 0,
    results: [],
  };

  function setTheme(theme) {
    var root = document.documentElement;
    root.classList.toggle('dark', theme === 'dark');
    try {
      localStorage.setItem('theme', theme);
    } catch (_) {}
  }

  function initTheme() {
    try {
      var saved = localStorage.getItem('theme');
      if (saved === 'dark' || saved === 'light') {
        setTheme(saved);
        return;
      }
    } catch (_) {}

    var prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    setTheme(prefersDark ? 'dark' : 'light');
  }

  function initThemeToggles() {
    document.querySelectorAll('[data-theme-toggle]').forEach(function (button) {
      button.addEventListener('click', function () {
        var root = document.documentElement;
        setTheme(root.classList.contains('dark') ? 'light' : 'dark');
      });
    });
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
      var sections = Array.prototype.slice.call(root.querySelectorAll('#installation, #usage, #examples, #api-reference, #view-source'));
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

  function initHighlighting() {
    function runHighlight() {
      if (window.hljs && typeof window.hljs.highlightAll === 'function') {
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
