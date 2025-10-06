<template>
  <div class="chart-container">
    <div class="chart-wrapper">
      <div class="axis-label label-top">Strong</div>
      <div class="axis-label label-bottom">Weak</div>
      <div class="axis-label label-left">Niche</div>
      <div class="axis-label label-right">Popular</div>
      
      <svg viewBox="0 0 560 500" class="chart-svg">
        <!-- Grid lines -->
        <line
          v-for="i in 7"
          :key="`h-${i}`"
          :x1="margin.left"
          :y1="margin.top + (height / 6) * (i - 1)"
          :x2="margin.left + width"
          :y2="margin.top + (height / 6) * (i - 1)"
          class="grid-line"
        />
        <line
          v-for="i in 7"
          :key="`v-${i}`"
          :x1="margin.left + (width / 6) * (i - 1)"
          :y1="margin.top"
          :x2="margin.left + (width / 6) * (i - 1)"
          :y2="margin.top + height"
          class="grid-line"
        />

        <!-- Center axes -->
        <line
          :x1="margin.left"
          :y1="centerY"
          :x2="margin.left + width"
          :y2="centerY"
          class="axis-line"
        />
        <line
          :x1="centerX"
          :y1="margin.top"
          :x2="centerX"
          :y2="margin.top + height"
          class="axis-line"
        />

        <!-- Data points -->
        <g v-for="(hero, i) in data" :key="hero.Name">
          <!-- <text
            :x="scaleX(hero.PickRate)"
            :y="scaleY(hero.WinRate) - 15"
            text-anchor="middle"
            fill="#fff"
            font-size="12"
            font-weight="600"
          >
            {{ hero.Name }}
          </text> -->
          <circle
            :cx="scaleX(hero.PickRate)"
            :cy="scaleY(hero.WinRate)"
            r="10"
            :fill="hero.Color"
            opacity="0.75"
            class="hero-dot"
            @mouseenter="handleMouseEnter(hero, $event)"
            @mousemove="handleMouseMove($event)"
            @mouseleave="handleMouseLeave"
          />
        </g>
      </svg>

      <div
        v-if="tooltip.show"
        class="tooltip"
        :style="{
          left: `${tooltip.x + 15}px`,
          top: `${tooltip.y - 15}px`
        }"
      >
        <img 
          class="tooltip-image"
          :src="tooltip.data.Image"
        />
        <div class="tooltip-name">{{ tooltip.data.Name }}</div>
        <div class="tooltip-stat">Pick Rate: {{ tooltip.data.PickRate }}%</div>
        <div class="tooltip-stat">Win Rate: {{ tooltip.data.WinRate }}%</div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'OverwatchScatterPlot',
  props: {
    heroData: Object,
  },
  data() {
    return {
      data: this.heroData,
      margin: { top: 40, right: 60, bottom: 30, left: 60 },
      tooltip: {
        show: false,
        x: 0,
        y: 0,
        data: null
      }
    };
  },
  watch: {
    heroData: {
      handler(newVal) {
        this.data = newVal;
      },
      deep: true,
      immediate: true
    }
  },
  mounted() {
    this.$el.addEventListener('hero-data-updated', (event) => {
      // This updates internal data, not the prop
      this.data = event.detail;
    });
  },
  computed: {
    width() {
      return 560 - this.margin.left - this.margin.right;
    },
    height() {
      return 500 - this.margin.top - this.margin.bottom;
    },
    pickRates() {
      return this.data.map(d => d.PickRate);
    },
    winRates() {
      return this.data.map(d => d.WinRate);
    },
    minPickRate() {
      return Math.min(...this.pickRates);
    },
    maxPickRate() {
      return Math.max(...this.pickRates);
    },
    minWinRate() {
      return Math.min(...this.winRates);
    },
    maxWinRate() {
      return Math.max(...this.winRates);
    },
    centerPickRate() {
      return (this.minPickRate + this.maxPickRate) / 2;
    },
    centerWinRate() {
      return (this.minWinRate + this.maxWinRate) / 2;
    },
    pickRateRange() {
      return this.maxPickRate - this.minPickRate;
    },
    winRateRange() {
      return this.maxWinRate - this.minWinRate;
    },
    paddedMinPickRate() {
      return this.minPickRate - this.pickRateRange * 0.1;
    },
    paddedMaxPickRate() {
      return this.maxPickRate + this.pickRateRange * 0.1;
    },
    paddedMinWinRate() {
      return this.minWinRate - this.winRateRange * 0.1;
    },
    paddedMaxWinRate() {
      return this.maxWinRate + this.winRateRange * 0.1;
    },
    centerX() {
      return this.scaleX(this.centerPickRate);
    },
    centerY() {
      return this.scaleY(this.centerWinRate);
    }
  },
  methods: {
    scaleX(value) {
      return this.margin.left + (value - this.paddedMinPickRate) / 
             (this.paddedMaxPickRate - this.paddedMinPickRate) * this.width;
    },
    scaleY(value) {
      return this.margin.top + this.height - (value - this.paddedMinWinRate) / 
             (this.paddedMaxWinRate - this.paddedMinWinRate) * this.height;
    },
    handleMouseEnter(hero, event) {
      this.tooltip = {
        show: true,
        x: event.pageX,
        y: event.pageY,
        data: hero
      };
    },
    handleMouseMove(event) {
      if (this.tooltip.show) {
        this.tooltip.x = event.pageX;
        this.tooltip.y = event.pageY;
      }
    },
    handleMouseLeave() {
      this.tooltip = {
        show: false,
        x: 0,
        y: 0,
        data: null
      };
    }
  }
};
</script>

<style scoped>

.chart-wrapper {
  position: relative;
  width: 560px;
  height: 500px;
  border-radius: 10px;
}

.axis-label {
  position: absolute;
  font-size: 20px;
  font-weight: 600;
  color: #fbbf24;
  padding: 4px 8px;
  border-radius: 4px;
}

.label-top {
  top: 0;
  left: 50%;
  transform: translateX(-50%);
}

.label-bottom {
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
}

.label-left {
  top: 50%;
  left: -15px;
  transform: translateY(-50%);
}

.label-right {
  top: 50%;
  right: -30px;
  transform: translateY(-50%);
}

.chart-svg {
  width: 100%;
  height: 100%;
}

.grid-line {
  stroke: rgba(255, 255, 255, 0.1);
  stroke-width: 1;
}

.axis-line {
  stroke: rgba(255, 255, 255, 0.5);
  stroke-width: 2;
}

.hero-dot {
  cursor: pointer;
  transition: all 0.3s ease;
}

.hero-dot:hover {
  opacity: 1 !important;
}

.tooltip {
  position: fixed;
  background: rgba(0, 0, 0, 0.9);
  color: #fff;
  padding: 10px 15px;
  border-radius: 8px;
  font-size: 13px;
  pointer-events: none;
  z-index: 100;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.tooltip-image {
  display: block;
  width: 50px;
  height: 50px;
  object-fit: contain;
  margin-bottom: 5px;
}

.tooltip-name {
  font-weight: bold;
  font-size: 15px;
  margin-bottom: 5px;
  color: #fbbf24;
}

.tooltip-stat {
  margin: 3px 0;
}
</style>