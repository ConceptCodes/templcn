(function () {
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
    setTheme(window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
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

  document.addEventListener('DOMContentLoaded', function () {
    initTheme();
    initThemeToggles();
    initTabs();
    initDropdownMenus();
  });
})();
