(function () {
  var searchIndex = null;
  var isSearchLoaded = false;

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
    setTheme('light');
  }

  function initThemeToggles() {
    document.querySelectorAll('[data-theme-toggle]').forEach(function (button) {
      button.addEventListener('click', function () {
        var root = document.documentElement;
        setTheme(root.classList.contains('dark') ? 'light' : 'dark');
      });
    });
  }

  function activateTab(root, value) {
    root.querySelectorAll('[data-tabs-trigger]').forEach(function (button) {
      var active = button.dataset.value === value;
      button.dataset.state = active ? 'active' : 'inactive';
      button.setAttribute('aria-selected', active ? 'true' : 'false');
      button.tabIndex = active ? 0 : -1;
    });
    root.querySelectorAll('[data-tabs-content]').forEach(function (panel) {
      var active = panel.dataset.value === value;
      panel.dataset.state = active ? 'active' : 'inactive';
      panel.hidden = !active;
    });
  }

  function initTabs() {
    document.querySelectorAll('[data-tabs-root]').forEach(function (root) {
      var defaultTrigger = root.querySelector('[data-tabs-trigger][data-state="active"]');
      var defaultValue =
        (defaultTrigger && defaultTrigger.dataset.value) ||
        root.dataset.value ||
        root.dataset.defaultValue ||
        (root.querySelector('[data-tabs-trigger]') && root.querySelector('[data-tabs-trigger]').dataset.value);
      if (defaultValue) {
        activateTab(root, defaultValue);
      }
      root.querySelectorAll('[data-tabs-trigger]').forEach(function (button) {
        button.addEventListener('click', function () {
          activateTab(root, button.dataset.value);
        });
      });
    });
  }

  function initDropdownMenus() {
    document.querySelectorAll('[data-dropdown-menu-root]').forEach(function (root) {
      var trigger = root.querySelector('[data-dropdown-menu-trigger]');
      var content = root.querySelector('[data-dropdown-menu-content]');
      if (!trigger || !content) return;

      function closeMenu() {
        root.dataset.open = 'false';
        content.hidden = true;
      }

      function openMenu() {
        root.dataset.open = 'true';
        content.hidden = false;
      }

      trigger.addEventListener('click', function (event) {
        event.preventDefault();
        if (content.hidden) openMenu(); else closeMenu();
      });

      document.addEventListener('click', function (event) {
        if (!root.contains(event.target)) closeMenu();
      });

      document.addEventListener('keydown', function (event) {
        if (event.key === 'Escape') closeMenu();
      });
    });
  }

  function loadSearchIndex() {
    if (isSearchLoaded) return Promise.resolve(searchIndex);
    return fetch('/search-index.json')
      .then(function (response) {
        if (!response.ok) throw new Error('Failed to load search index');
        return response.json();
      })
      .then(function (data) {
        searchIndex = data;
        isSearchLoaded = true;
        return searchIndex;
      })
      .catch(function (error) {
        console.error('Error loading search index:', error);
        return [];
      });
  }

  function showSearchDialog() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    if (dialog && input) {
      dialog.classList.remove('hidden');
      input.focus();
    }
  }

  function hideSearchDialog() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    if (dialog) {
      dialog.classList.add('hidden');
    }
    if (input) {
      input.value = '';
    }
    var resultsContainer = document.querySelector('[data-search-results="true"]');
    if (resultsContainer) {
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">Type to search...</p>';
    }
  }

  function renderSearchResults(query) {
    var resultsContainer = document.querySelector('[data-search-results="true"]');
    if (!resultsContainer) return;

    if (!query || query.trim() === '') {
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">Type to search...</p>';
      return;
    }

    var lowerQuery = query.toLowerCase();
    var matches = searchIndex.filter(function (entry) {
      return entry.title.toLowerCase().indexOf(lowerQuery) !== -1;
    });

    if (matches.length === 0) {
      resultsContainer.innerHTML = '<p class="py-4 text-center text-sm text-muted-foreground">No results found</p>';
      return;
    }

    var html = '<ul class="flex flex-col gap-1">';
    matches.forEach(function (entry) {
      var typeClass = '';
      var typeLabel = '';
      if (entry.type === 'component') {
        typeClass = 'bg-primary/10 text-primary';
        typeLabel = 'Component';
      } else if (entry.type === 'docs') {
        typeClass = 'bg-secondary/10 text-secondary-foreground';
        typeLabel = 'Docs';
      } else if (entry.type === 'chart') {
        typeClass = 'bg-accent/10 text-accent-foreground';
        typeLabel = 'Chart';
      }

      html += '<li><a href="' + entry.url + '" class="flex items-center justify-between gap-4 rounded-md px-3 py-2 text-sm hover:bg-accent hover:text-accent-foreground">';
      html += '<span>' + entry.title + '</span>';
      html += '<span class="rounded px-1.5 py-0.5 text-xs ' + typeClass + '">' + typeLabel + '</span>';
      html += '</a></li>';
    });
    html += '</ul>';
    resultsContainer.innerHTML = html;
  }

  function initSearch() {
    var dialog = document.querySelector('[data-search-dialog="true"]');
    var backdrop = document.querySelector('[data-search-backdrop="true"]');
    var input = document.querySelector('[data-search-input="true"]');
    var searchResults = document.querySelector('[data-search-results="true"]');

    if (!dialog || !input) return;

    document.addEventListener('keydown', function (event) {
      if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
        event.preventDefault();
        if (dialog.classList.contains('hidden')) {
          showSearchDialog();
        } else {
          hideSearchDialog();
        }
      } else if (event.key === 'Escape' && !dialog.classList.contains('hidden')) {
        hideSearchDialog();
      }
    });

    if (backdrop) {
      backdrop.addEventListener('click', hideSearchDialog);
    }

    if (input) {
      input.addEventListener('input', function (event) {
        var query = event.target.value;
        loadSearchIndex().then(function () {
          renderSearchResults(query);
        });
      });
    }

    if (searchResults) {
      searchResults.addEventListener('click', function (event) {
        var link = event.target.closest('a');
        if (link) {
          hideSearchDialog();
        }
      });
    }
  }

  document.addEventListener('DOMContentLoaded', function () {
    initTheme();
    initThemeToggles();
    initTabs();
    initDropdownMenus();
    initMobileMenu();
    initSearch();
  });

  function initMobileMenu() {
    var toggle = document.querySelector('[data-mobile-menu-toggle]');
    var drawer = document.querySelector('[data-mobile-drawer]');
    var close = document.querySelector('[data-mobile-drawer-close]');
    if (toggle && drawer) {
      toggle.addEventListener('click', function () {
        drawer.classList.toggle('hidden');
      });
      if (close) {
        close.addEventListener('click', function () {
          drawer.classList.add('hidden');
        });
      }
      drawer.querySelectorAll('a').forEach(function (link) {
        link.addEventListener('click', function () {
          drawer.classList.add('hidden');
        });
      });
    }
  }
})();
