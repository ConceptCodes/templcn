(function () {
  function showTooltip(root, content, x, y) {
    var tip = d3.select(root).selectAll('.chart-tooltip').data([0]);
    tip = tip.enter().append('div')
      .attr('class', 'chart-tooltip absolute z-50 rounded-lg border border-border/70 bg-popover px-3 py-2 text-xs text-popover-foreground shadow-md pointer-events-none opacity-0 transition-opacity')
      .merge(tip);
    
    tip.html(content)
      .style('left', (x + 10) + 'px')
      .style('top', (y - 10) + 'px')
      .style('opacity', 1);
  }

  function hideTooltip(root) {
    d3.select(root).selectAll('.chart-tooltip').style('opacity', 0);
  }

  function readPalette() {
    var styles = getComputedStyle(document.documentElement);
    var palette = ['--chart-1', '--chart-2', '--chart-3', '--chart-4', '--chart-5']
      .map(function (name) {
        return styles.getPropertyValue(name).trim();
      })
      .filter(Boolean);

    if (palette.length === 0) {
      palette = ['hsl(222 47% 11%)', 'hsl(173 58% 39%)', 'hsl(24 95% 53%)', 'hsl(269 81% 57%)', 'hsl(47 96% 53%)'];
    }

    return palette;
  }

  function getSeries(cfg) {
    var palette = readPalette();
    return (cfg.series || []).map(function (series, index) {
      return {
        key: series.key,
        label: series.label || series.key,
        color: series.color || palette[index % palette.length],
      };
    });
  }

  function getValue(row, key) {
    var value = Number(row[key]);
    return Number.isFinite(value) ? value : 0;
  }

  function createSvg(root, width, height) {
    return d3.select(root).append('svg').attr('viewBox', '0 0 ' + width + ' ' + height).attr('class', 'h-full w-full');
  }

  function renderGrid(svg, x1, x2, yScale, width, height, margin) {
    var ticks = yScale.ticks(4);
    var grid = svg.append('g').attr('stroke', 'currentColor').attr('stroke-opacity', 0.08);

    grid
      .selectAll('line')
      .data(ticks)
      .enter()
      .append('line')
      .attr('x1', x1)
      .attr('x2', x2)
      .attr('y1', function (d) {
        return yScale(d);
      })
      .attr('y2', function (d) {
        return yScale(d);
      });

    svg
      .append('g')
      .attr('fill', 'currentColor')
      .attr('fill-opacity', 0.45)
      .attr('class', 'text-[10px] font-medium')
      .selectAll('text')
      .data(ticks)
      .enter()
      .append('text')
      .attr('x', margin.left - 8)
      .attr('y', function (d) {
        return yScale(d) + 3;
      })
      .attr('text-anchor', 'end')
      .text(function (d) {
        return d;
      });
  }

  function renderAxes(svg, xScale, labels, width, height, margin) {
    svg
      .append('g')
      .attr('fill', 'currentColor')
      .attr('fill-opacity', 0.45)
      .attr('class', 'text-[10px] font-medium')
      .selectAll('text')
      .data(labels)
      .enter()
      .append('text')
      .attr('x', function (label) {
        return xScale(label);
      })
      .attr('y', height - margin.bottom + 18)
      .attr('text-anchor', 'middle')
      .text(function (label) {
        return label;
      });
  }

  function renderArea(root, cfg) {
    var width = Math.max(root.clientWidth, 280);
    var height = 288;
    var margin = { top: 16, right: 16, bottom: 36, left: 36 };
    var data = cfg.data || [];
    var series = getSeries(cfg);
    if (!data.length || !series.length) return;

    var labels = data.map(function (row) { return row.label; });
    var x = d3.scalePoint().domain(labels).range([margin.left, width - margin.right]).padding(0.5);
    
    var values = series.map(function (entry) {
      return data.map(function (row) { return getValue(row, entry.key); });
    });
    var yMax = d3.max(values.flat()) || 100;
    var stackedMax = null;

    if (cfg.stacked) {
      stackedMax = d3.max(data, function (row) {
        return d3.sum(series, function (entry) { return getValue(row, entry.key); });
      }) || 100;
    }

    var y = d3.scaleLinear().domain([0, (stackedMax || yMax) * 1.15]).nice().range([height - margin.bottom, margin.top]);

    var svg = createSvg(root, width, height);
    renderGrid(svg, margin.left, width - margin.right, y, width, height, margin);
    renderAxes(svg, x, labels, width, height, margin);

    if (cfg.stacked) {
      var stack = d3.stack().keys(series.map(function (entry) { return entry.key; }))(data);
      var area = d3.area()
        .x(function (point) { return x(point.data.label); })
        .y0(function (point) { return y(point[0]); })
        .y1(function (point) { return y(point[1]); })
        .curve(d3.curveMonotoneX);

      stack.forEach(function (layer, index) {
        svg.append('path')
          .datum(layer)
          .attr('d', area)
          .attr('fill', series[index].color)
          .attr('fill-opacity', 0.25)
          .on('mousemove', function(e, d) {
            var mouseX = d3.pointer(e)[0];
            var bisect = d3.bisector(function(d) { return x(d.label); }).center;
            var idx = bisect(data, mouseX);
            var row = data[idx];
            var content = '<strong>' + row.label + '</strong><br/>' + series[index].label + ': ' + getValue(row, series[index].key);
            showTooltip(root, content, e.offsetX, e.offsetY);
          })
          .on('mouseleave', function() { hideTooltip(root); });
      });
    } else {
      series.forEach(function (entry, index) {
        var line = d3.line()
          .x(function (row) { return x(row.label); })
          .y(function (row) { return y(getValue(row, entry.key)); })
          .curve(d3.curveMonotoneX);

        var area = d3.area()
          .x(function (row) { return x(row.label); })
          .y0(y(0))
          .y1(function (row) { return y(getValue(row, entry.key)); })
          .curve(d3.curveMonotoneX);

        svg.append('path')
          .datum(data)
          .attr('d', area)
          .attr('fill', entry.color)
          .attr('fill-opacity', 0.15);

        svg.append('path')
          .datum(data)
          .attr('d', line)
          .attr('fill', 'none')
          .attr('stroke', entry.color)
          .attr('stroke-width', 3);
      });
    }

    if (cfg.showPoints) {
      svg.append('g')
        .selectAll('circle')
        .data(data)
        .enter()
        .append('circle')
        .attr('cx', function (row) { return x(row.label); })
        .attr('cy', function (row) { return y(getValue(row, series[0].key)); })
        .attr('r', 4.5)
        .attr('fill', series[0].color)
        .attr('stroke', 'var(--background)')
        .attr('stroke-width', 2)
        .on('mousemove', function(e, row) {
          var content = '<strong>' + row.label + '</strong><br/>' + series[0].label + ': ' + getValue(row, series[0].key);
          showTooltip(root, content, e.offsetX, e.offsetY);
          d3.select(this).attr('r', 6);
        })
        .on('mouseleave', function() {
          hideTooltip(root);
          d3.select(this).attr('r', 4.5);
        });
    }
  }

  function renderBar(root, cfg) {
    var width = Math.max(root.clientWidth, 280);
    var height = 288;
    var margin = { top: 16, right: 16, bottom: 36, left: 44 };
    var data = cfg.data || [];
    var series = getSeries(cfg);
    if (!data.length || !series.length) return;

    var svg = createSvg(root, width, height);
    var labels = data.map(function (row) { return row.label; });
    var vertical = cfg.orientation !== 'horizontal';

    if (vertical) {
      var x0 = d3.scaleBand().domain(labels).range([margin.left, width - margin.right]).paddingInner(0.24);
      var x1 = d3.scaleBand().domain(series.map(function (entry) { return entry.key; })).range([0, x0.bandwidth()]).padding(0.14);
      var y = d3.scaleLinear().domain([0, d3.max(data, function (row) {
        return d3.max(series, function (entry) { return getValue(row, entry.key); });
      }) * 1.15 || 100]).nice().range([height - margin.bottom, margin.top]);

      renderGrid(svg, margin.left, width - margin.right, y, width, height, margin);

      svg.append('g')
        .selectAll('g')
        .data(data)
        .enter()
        .append('g')
        .attr('transform', function (row) { return 'translate(' + x0(row.label) + ',0)'; })
        .selectAll('rect')
        .data(function (row) { return series.map(function (entry) { return { series: entry, value: getValue(row, entry.key), label: row.label }; }); })
        .enter()
        .append('rect')
        .attr('x', function (entry) { return x1(entry.series.key); })
        .attr('y', function (entry) { return y(entry.value); })
        .attr('width', x1.bandwidth())
        .attr('height', function (entry) { return y(0) - y(entry.value); })
        .attr('rx', 4)
        .attr('fill', function (entry) { return entry.series.color; })
        .on('mousemove', function(e, d) {
          var content = '<strong>' + d.label + '</strong><br/>' + d.series.label + ': ' + d.value;
          showTooltip(root, content, e.offsetX, e.offsetY);
          d3.select(this).attr('fill-opacity', 0.8);
        })
        .on('mouseleave', function() {
          hideTooltip(root);
          d3.select(this).attr('fill-opacity', 1);
        });

      renderAxes(svg, function (label) { return x0(label) + x0.bandwidth() / 2; }, labels, width, height, margin);
    } else {
      var y0 = d3.scaleBand().domain(labels).range([margin.top, height - margin.bottom]).paddingInner(0.24);
      var y1 = d3.scaleBand().domain(series.map(function (entry) { return entry.key; })).range([0, y0.bandwidth()]).padding(0.14);
      var x = d3.scaleLinear().domain([0, d3.max(data, function (row) {
        return d3.max(series, function (entry) { return getValue(row, entry.key); });
      }) * 1.15 || 100]).nice().range([margin.left, width - margin.right]);

      var grid = svg.append('g').attr('stroke', 'currentColor').attr('stroke-opacity', 0.08);
      x.ticks(4).forEach(function (tick) {
        grid.append('line').attr('x1', x(tick)).attr('x2', x(tick)).attr('y1', margin.top).attr('y2', height - margin.bottom);
      });

      svg.append('g')
        .selectAll('g')
        .data(data)
        .enter()
        .append('g')
        .attr('transform', function (row) { return 'translate(0,' + y0(row.label) + ')'; })
        .selectAll('rect')
        .data(function (row) { return series.map(function (entry) { return { series: entry, value: getValue(row, entry.key), label: row.label }; }); })
        .enter()
        .append('rect')
        .attr('x', margin.left)
        .attr('y', function (entry) { return y1(entry.series.key); })
        .attr('width', function (entry) { return x(entry.value) - x(0); })
        .attr('height', y1.bandwidth())
        .attr('rx', 4)
        .attr('fill', function (entry) { return entry.series.color; })
        .on('mousemove', function(e, d) {
          var content = '<strong>' + d.label + '</strong><br/>' + d.series.label + ': ' + d.value;
          showTooltip(root, content, e.offsetX, e.offsetY);
          d3.select(this).attr('fill-opacity', 0.8);
        })
        .on('mouseleave', function() {
          hideTooltip(root);
          d3.select(this).attr('fill-opacity', 1);
        });

      svg.append('g')
        .attr('fill', 'currentColor')
        .attr('fill-opacity', 0.45)
        .attr('class', 'text-[10px] font-medium')
        .selectAll('text')
        .data(labels)
        .enter()
        .append('text')
        .attr('x', margin.left - 8)
        .attr('y', function (label) { return y0(label) + y0.bandwidth() / 2 + 3; })
        .attr('text-anchor', 'end')
        .text(function (label) { return label; });
    }
  }

  function renderLine(root, cfg) {
    var width = Math.max(root.clientWidth, 280);
    var height = 288;
    var margin = { top: 16, right: 16, bottom: 36, left: 36 };
    var data = cfg.data || [];
    var series = getSeries(cfg);
    if (!data.length || !series.length) return;

    var labels = data.map(function (row) { return row.label; });
    var x = d3.scalePoint().domain(labels).range([margin.left, width - margin.right]).padding(0.5);
    var yMax = d3.max(series, function (entry) {
      return d3.max(data, function (row) { return getValue(row, entry.key); });
    }) || 100;
    var y = d3.scaleLinear().domain([0, yMax * 1.15]).nice().range([height - margin.bottom, margin.top]);
    var svg = createSvg(root, width, height);

    renderGrid(svg, margin.left, width - margin.right, y, width, height, margin);
    renderAxes(svg, x, labels, width, height, margin);

    series.forEach(function (entry, index) {
      var line = d3.line()
        .x(function (row) { return x(row.label); })
        .y(function (row) { return y(getValue(row, entry.key)); })
        .curve(cfg.step ? d3.curveStepAfter : d3.curveMonotoneX);

      svg.append('path')
        .datum(data)
        .attr('d', line)
        .attr('fill', 'none')
        .attr('stroke', entry.color)
        .attr('stroke-width', 3);

      if (cfg.showPoints !== false) {
        svg.append('g')
          .selectAll('circle')
          .data(data)
          .enter()
          .append('circle')
          .attr('cx', function (row) { return x(row.label); })
          .attr('cy', function (row) { return y(getValue(row, entry.key)); })
          .attr('r', 4.5)
          .attr('fill', entry.color)
          .attr('stroke', 'var(--background)')
          .attr('stroke-width', 2)
          .on('mousemove', function(e, row) {
            var content = '<strong>' + row.label + '</strong><br/>' + entry.label + ': ' + getValue(row, entry.key);
            showTooltip(root, content, e.offsetX, e.offsetY);
            d3.select(this).attr('r', 6.5);
          })
          .on('mouseleave', function() {
            hideTooltip(root);
            d3.select(this).attr('r', 4.5);
          });
      }
    });
  }

  function renderPie(root, cfg) {
    var width = Math.max(root.clientWidth, 280);
    var height = 288;
    var radius = Math.min(width, height) / 2 - 24;
    var data = cfg.data || [];
    if (!data.length) return;

    var palette = readPalette();
    var svg = createSvg(root, width, height);
    var g = svg.append('g').attr('transform', 'translate(' + width / 2 + ',' + height / 2 + ')');
    
    var pie = d3.pie().value(function (d) { return Number(d.value) || 0; }).sort(null);
    var innerRadius = Math.max(0, Math.min(radius * (Number(cfg.innerRadius) || 0), radius - 12));
    var arc = d3.arc().innerRadius(innerRadius).outerRadius(radius);
    var hoverArc = d3.arc().innerRadius(innerRadius).outerRadius(radius + 6);

    g.selectAll('path')
      .data(pie(data))
      .enter()
      .append('path')
      .attr('d', arc)
      .attr('fill', function (d, index) { return data[index].color || palette[index % palette.length]; })
      .attr('stroke', 'var(--background)')
      .attr('stroke-width', 2)
      .on('mousemove', function(e, d) {
        var content = '<strong>' + d.data.label + '</strong><br/>Value: ' + d.data.value;
        showTooltip(root, content, e.offsetX, e.offsetY);
        d3.select(this).transition().duration(200).attr('d', hoverArc);
      })
      .on('mouseleave', function() {
        hideTooltip(root);
        d3.select(this).transition().duration(200).attr('d', arc);
      });

    if (cfg.centerLabel) {
      g.append('text').attr('text-anchor', 'middle').attr('dy', '-0.1em').attr('class', 'fill-foreground text-4xl font-semibold').text(cfg.centerLabel);
    }
    if (cfg.centerCaption) {
      g.append('text').attr('text-anchor', 'middle').attr('dy', '1.4em').attr('class', 'fill-muted-foreground text-sm font-medium').text(cfg.centerCaption);
    }
  }

  function renderRadial(root, cfg) {
    var width = Math.max(root.clientWidth, 280);
    var height = 288;
    var radius = Math.min(width, height) / 2 - 24;
    var data = cfg.segments || [];
    var value = Number(cfg.value) || 0;
    var max = Number(cfg.max) || 100;
    var palette = readPalette();
    var svg = createSvg(root, width, height);
    var g = svg.append('g').attr('transform', 'translate(' + width / 2 + ',' + height / 2 + ')');

    if (data.length > 0) {
      var pie = d3.pie().value(function (d) { return Number(d.value) || 0; }).sort(null);
      var arc = d3.arc().innerRadius(radius - 28).outerRadius(radius);
      var hoverArc = d3.arc().innerRadius(radius - 32).outerRadius(radius + 4);

      g.selectAll('path')
        .data(pie(data))
        .enter()
        .append('path')
        .attr('d', arc)
        .attr('fill', function (d, index) { return data[index].color || palette[index % palette.length]; })
        .attr('stroke', 'var(--background)')
        .attr('stroke-width', 2)
        .on('mousemove', function(e, d) {
          var content = '<strong>' + d.data.label + '</strong><br/>' + d.data.value;
          showTooltip(root, content, e.offsetX, e.offsetY);
          d3.select(this).transition().duration(200).attr('d', hoverArc);
        })
        .on('mouseleave', function() {
          hideTooltip(root);
          d3.select(this).transition().duration(200).attr('d', arc);
        });
    } else {
      var track = d3.arc().innerRadius(radius - 28).outerRadius(radius);
      var progress = d3.arc().innerRadius(radius - 28).outerRadius(radius);
      g.append('path').attr('d', track({ startAngle: 0, endAngle: Math.PI * 2 })).attr('fill', 'currentColor').attr('fill-opacity', 0.08);
      g.append('path')
        .attr('d', progress({
          startAngle: -Math.PI / 2,
          endAngle: -Math.PI / 2 + (Math.PI * 2 * Math.max(0, Math.min(value, max))) / max,
        }))
        .attr('fill', palette[0])
        .on('mousemove', function(e) {
          showTooltip(root, '<strong>Progress</strong><br/>' + value + '/' + max, e.offsetX, e.offsetY);
        })
        .on('mouseleave', function() { hideTooltip(root); });
    }

    if (cfg.centerLabel) {
      g.append('text').attr('text-anchor', 'middle').attr('dy', '-0.1em').attr('class', 'fill-foreground text-4xl font-semibold').text(cfg.centerLabel);
    }
    if (cfg.centerCaption) {
      g.append('text').attr('text-anchor', 'middle').attr('dy', '1.4em').attr('class', 'fill-muted-foreground text-xs uppercase tracking-[0.24em] font-bold').text(cfg.centerCaption);
    }
  }

  function renderChart(root) {
    var type = root.dataset.chart;
    var configRaw = root.dataset.chartConfig || '{}';
    var cfg;
    try {
      cfg = JSON.parse(configRaw);
    } catch (_) {
      cfg = {};
    }

    root.innerHTML = '';
    root.classList.add('overflow-hidden');

    if (type === 'area') renderArea(root, cfg);
    else if (type === 'bar') renderBar(root, cfg);
    else if (type === 'line') renderLine(root, cfg);
    else if (type === 'pie') renderPie(root, cfg);
    else if (type === 'radial') renderRadial(root, cfg);
  }

  function initCharts() {
    var nodes = document.querySelectorAll('[data-chart]');
    if (!nodes.length || !window.d3) return;

    nodes.forEach(function (node) {
      var observer = new ResizeObserver(function () {
        renderChart(node);
      });

      observer.observe(node);
      renderChart(node);
      node.__chartObserver = observer;
    });
  }

  document.addEventListener('DOMContentLoaded', initCharts);
})();
