<template>
  <div class="chart-container">
    <div class="chart-wrapper">
      <div class="axis-label label-top">Strong</div>
      <div class="axis-label label-bottom">Weak</div>
      <div class="axis-label label-left">Niche</div>
      <div class="axis-label label-right">Popular</div>
      
      <svg viewBox="0 0 560 500" class="chart-svg" preserveAspectRatio="xMidYMid meet">
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
        <g v-for="(hero, i) in data" :key="getItemKey(hero)">
          <defs v-if="hero.Heroes.length > 1">
            <radialGradient :id="`combo-gradient-${getItemKey(hero)}`">
              <template v-for="(h, index) in hero.Heroes" :key="`${h.Name}-stops`">
                <stop 
                  :offset="`${(index / hero.Heroes.length) * 100}%`" 
                  :stop-color="h.Color" 
                />
                <stop 
                  :offset="`${((index + 1) / hero.Heroes.length) * 100}%`" 
                  :stop-color="h.Color" 
                />
              </template>
            </radialGradient>
          </defs>
          <text
            v-if="tooltip.pinnedHero !== getDisplayName(hero) && (tooltip.pinned || hoveredHero !== getDisplayName(hero))"
            :x="getLabelPosition(hero, i).x"
            :y="getLabelPosition(hero, i).y"
            text-anchor="middle"
            fill="#fff"
            font-size="12"
            font-weight="600"
            class="hero-dot-text"
          >
            {{ getDisplayName(hero) }}
          </text>
          <circle
            :cx="scaleX(hero.Stats.PickRate)"
            :cy="scaleY(hero.Stats.WinRate)"
            r="10"
            :fill="hero.Heroes.length === 1 ? hero.Heroes[0].Color : `url(#combo-gradient-${getItemKey(hero)})`"
            opacity="0.75"
            :class="{ 'hero-dot': true, 'pinned': tooltip.pinned && tooltip.pinnedHero === getDisplayName(hero) }"
            @mouseenter="handleMouseEnter(hero, $event)"
            @mousemove="handleMouseMove($event)"
            @mouseleave="handleMouseLeave"
            @click="handleDotClick(hero, $event)"
          />
          <circle
            v-if="tooltip.pinned && tooltip.pinnedHero === getDisplayName(hero) || hoveredHero === getDisplayName(hero)"
            :cx="scaleX(hero.Stats.PickRate)"
            :cy="scaleY(hero.Stats.WinRate)"
            r="10"
            fill="none"
            :stroke="hero.Heroes[0].Color"
            opacity="1"
            stroke-width="3"
            class="pulse-ring"
          />

        </g>
      </svg>

      <div
        v-if="tooltip.show"
        class="tooltip"
        :style="{
          left: `${tooltip.x + 5}px`,
          top: `${tooltip.y - 5}px`
        }"
      >
        <button 
          v-if="tooltip.pinned"
          class="tooltip-close"
          @click="closeTooltip"
        >
          ×
        </button>
        <img 
            v-if="tooltip.data.Heroes.length === 1"
            class="tooltip-image"
            :src="tooltip.data.Heroes[0].Image"
        />
        <template v-else>
            <img 
                v-for="hero in tooltip.data.Heroes"
                :key="hero.ID"
                class="tooltip-image combo-image"
                :src="hero.Image"
            />
        </template>
        <div class="tooltip-name">{{ getDisplayName(tooltip.data) }}</div>
        <div class="tooltip-stat">Pick Rate: {{ tooltip.data.Stats.PickRate }}%</div>
        <div class="tooltip-stat">Win Rate: {{ tooltip.data.Stats.WinRate }}%</div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'DeadlockScatterPlot',
  props: {
    heroData: Object,
  },
  data() {
    return {
      data: this.heroData,
      margin: { top: 40, right: 60, bottom: 30, left: 60 },
      hoveredHero: null, 
      tooltip: {
        show: false,
        pinned: false,
        pinnedHero: null,
        x: 0,
        y: 0,
        data: null
      },
      labelPositions: []
    };
  },
  watch: {
    heroData: {
      handler(newVal) {
        this.data = newVal;
      },
      deep: true,
      immediate: true
    },
    data: {
      handler() {
        // Reset tooltip whenever data changes
        this.tooltip = {
          show: false,
          pinned: false,
          pinnedHero: null,
          x: 0,
          y: 0,
          data: null
        };
        this.hoveredHero = null;
        if (this.data && this.data.length > 0) {
          this.$nextTick(() => {
            this.calculateLabelPositions();
          });
        }
      },
      deep: true
    }
  },
  mounted() {
    this.$el.addEventListener('hero-data-updated', (event) => {
      // This updates internal data, not the prop
      this.data = event.detail;
    });

    // Wait for data to be fully set before calculating positions
    this.$nextTick(() => {
      if (this.data && this.data.length > 0) {
        this.calculateLabelPositions();
      }
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
      return this.data.map(d => d.Stats.PickRate);
    },
    winRates() {
      return this.data.map(d => d.Stats.WinRate);
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
    calculateLabelPositions() {
      this.labelPositions = [];

      const positions = [];
      const labelHeight = 16; // Approximate height of label with padding
      
      this.data.forEach((hero, i) => {
        const x = this.scaleX(hero.Stats.PickRate);
        const y = this.scaleY(hero.Stats.WinRate);
        
        // Estimate label width based on character count (roughly 7px per character)
        const nameLength = this.getDisplayName(hero).length;
        const estimatedWidth = nameLength * 7;
        const halfWidth = estimatedWidth / 2;
        
        // Possible offsets to try (top, bottom, left, right, diagonal)
        // Adjust horizontal offsets based on name length
        const offsets = [
          { x: 0, y: -12 },                    // top (default)
          { x: 0, y: 20 },                     // bottom
          { x: -halfWidth - 10, y: 4 },        // left (adjusted)
          { x: halfWidth + 10, y: 4 },         // right (adjusted)
          { x: -halfWidth - 5, y: -10 },       // top-left (adjusted)
          { x: halfWidth + 5, y: -10 },        // top-right (adjusted)
          { x: -halfWidth - 5, y: 20 },        // bottom-left (adjusted)
          { x: halfWidth + 5, y: 20 }          // bottom-right (adjusted)
        ];
        
        let bestOffset = offsets[0];
        let bestScore = -Infinity;
        
        // Try each offset and score based on distance from other labels AND dots
        for (const offset of offsets) {
          const testX = x + offset.x;
          const testY = y + offset.y;
          
          let minLabelDist = Infinity;
          let minDotDist = Infinity;
          
          // Check distance to all existing labels (rectangular collision)
          for (let j = 0; j < positions.length; j++) {
            const pos = positions[j];
            const otherHero = this.data[j];
            const otherWidth = this.getDisplayName(otherHero).length * 7;
            
            // Calculate rectangular overlap
            const xDist = Math.abs(testX - pos.x);
            const yDist = Math.abs(testY - pos.y);
            const widthSum = (estimatedWidth + otherWidth) / 2 + 10;
            
            // Penalize if labels would overlap
            if (xDist < widthSum && yDist < labelHeight) {
              minLabelDist = Math.min(minLabelDist, 0);
            } else {
              const dist = Math.sqrt(Math.pow(xDist, 2) + Math.pow(yDist, 2));
              minLabelDist = Math.min(minLabelDist, dist);
            }
          }
          
          // Check distance to all dots (avoid covering other dots)
          for (let j = 0; j < this.data.length; j++) {
            if (j === i) continue; // Skip own dot
            
            const otherHero = this.data[j];
            const dotX = this.scaleX(otherHero.Stats.PickRate);
            const dotY = this.scaleY(otherHero.Stats.WinRate);
            
            // Check if label would cover the dot
            const xDist = Math.abs(testX - dotX);
            const yDist = Math.abs(testY - dotY);
            
            if (xDist < halfWidth && yDist < labelHeight / 2) {
              minDotDist = Math.min(minDotDist, 0); // Label covers dot
            } else {
              const dist = Math.sqrt(Math.pow(xDist, 2) + Math.pow(yDist, 2));
              minDotDist = Math.min(minDotDist, dist);
            }
          }
          
          // Score this position (prioritize avoiding overlaps, then maximize distance)
          const score = minLabelDist * 2 + minDotDist;
          
          // If this position is better, use it
          if (score > bestScore) {
            bestScore = score;
            bestOffset = offset;
          }
        }
        
        positions.push({
          x: x + bestOffset.x,
          y: y + bestOffset.y,
          width: estimatedWidth
        });
      });
      
      this.labelPositions = positions;
    },
    getLabelPosition(hero, index) {
      if (this.labelPositions[index]) {
        return this.labelPositions[index];
      }
      // Fallback to default position
      return {
        x: this.scaleX(hero.Stats.PickRate),
        y: this.scaleY(hero.Stats.WinRate) - 12
      };
    },
    handleMouseEnter(hero, event) {
      this.hoveredHero = this.getDisplayName(hero);
      // Don't show tooltip on hover if it's pinned to a different hero
      if (this.tooltip.pinned && this.tooltip.pinnedHero !== this.getDisplayName(hero)) {
        return;
      }
      
      this.tooltip = {
        show: true,
        pinned: this.tooltip.pinned,
        pinnedHero: this.tooltip.pinnedHero,
        x: event.pageX,
        y: event.pageY,
        data: hero
      };
    },
    handleMouseMove(event) {
      // Don't update position if tooltip is pinned
      if (this.tooltip.show && !this.tooltip.pinned) {
        this.tooltip.x = event.pageX;
        this.tooltip.y = event.pageY;
      }
    },
    handleMouseLeave() {
      this.hoveredHero = null;
      // Don't hide tooltip if it's pinned
      if (!this.tooltip.pinned) {
        this.tooltip = {
          show: false,
          pinned: false,
          pinnedHero: null,
          x: 0,
          y: 0,
          data: null
        };
      }
    },
    handleDotClick(hero, event) {
      event.stopPropagation();
      
      // If clicking the same hero that's pinned, unpin it
      if (this.tooltip.pinned && this.tooltip.pinnedHero === this.getDisplayName(hero)) {
        this.tooltip = {
          show: false,
          pinned: false,
          pinnedHero: null,
          x: 0,
          y: 0,
          data: null
        };
      } else {
        // Pin the tooltip to this hero
        this.tooltip = {
          show: true,
          pinned: true,
          pinnedHero: this.getDisplayName(hero),
          x: event.pageX,
          y: event.pageY,
          data: hero
        };
      }
    },
    closeTooltip() {
      this.tooltip = {
        show: false,
        pinned: false,
        pinnedHero: null,
        x: 0,
        y: 0,
        data: null
      };
      this.hoveredHero = null;
    },
    getItemKey(item) {
      // console.log("Get Item Key Log")
      // console.log(item)
      if (item.Heroes.length === 1) {
        return item.Heroes[0].Name
      }

      const names = item.Heroes.map(h => h.Name);

      return names.join("-")
    },
    getDisplayName(item) {
        return item.Heroes.length === 1 ? item.Heroes[0].Name : `Combo ${this.getComboNumber(item)}`
    },
    getComboNumber(comboItem) {
        // Filter to get only combos from your data
        const combos = this.data.filter(item => item.Heroes.length > 1);
        // Find the index of this specific combo and add 1 (for human-readable numbering)
        return combos.findIndex(c => this.getComboIdentifier(c) === this.getComboIdentifier(comboItem)) + 1;
    },
    getComboIdentifier(item) {
        // Create unique identifier for combo (using hero names or IDs)
        return item.Heroes.map(h => h.ID).sort().join('-');
    }
  }
};
</script>

<style scoped>

.chart-container {
  width: 100%;
  max-width: 560px;
  margin: 0 auto;
  overflow: visible;
}

.chart-wrapper {
  position: relative;
  width: 100%;
  aspect-ratio: 560 / 500;
  border-radius: 10px;
  overflow: visible;
}

.axis-label {
  position: absolute;
  font-size: 20px;
  font-weight: 600;
  color: #fbbf24;
  padding: 4px 8px;
  border-radius: 4px;
  white-space: nowrap;
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
  left: -5px;
  transform: translateY(-50%);
}

.label-right {
  top: 50%;
  right: -20px;
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

.hero-dot-text {
  text-shadow:
    -1px -1px 0 #111827,
    1px -1px 0 #111827,
    -1px 1px 0 #111827,
    1px 1px 0 #111827;
  transition: all 0.3s ease;
}

.hero-dot {
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.hero-dot:hover {
  opacity: 1 !important;
}

.hero-dot.pinned {
  opacity: 1 !important;
}

.tooltip {
  position: fixed;
  background: rgba(0, 0, 0, 0.9);
  color: #fff;
  padding: 10px 15px;
  border-radius: 8px;
  font-size: 13px;
  pointer-events: auto;
  z-index: 100;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.tooltip-close {
  position: absolute;
  top: 5px;
  right: 5px;
  background: none;
  border: none;
  color: #fff;
  font-size: 24px;
  line-height: 1;
  cursor: pointer;
  padding: 0;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s ease;
}

.tooltip-close:hover {
  color: #fbbf24;
}

.tooltip-image {
  display: block;
  width: 50px;
  height: 50px;
  object-fit: contain;
  margin-bottom: 5px;
}

.combo-image {
  display: inline-block;
  width: 20px;
  height: 20px;
  margin-right: -8px;  /* Negative margin for overlap */
}

.combo-image:last-child {
  margin-right: 0;
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

.pulse-ring {
  animation: pulse-ring 2s ease-out infinite;
}

@keyframes pulse-ring {
  0% {
    r: 10;
    opacity: 1;
  }
  100% {
    r: 18;
    opacity: 0;
  }
}

/* Responsive adjustments */
@media (max-width: 640px) {
  .axis-label {
    font-size: 16px;
    padding: 2px 6px;
  }
  
  /* .hero-dot-text {
    font-size: 10px;
  } */
}

@media (max-width: 480px) {
  .axis-label {
    font-size: 14px;
    padding: 2px 4px;
  }
  
  /* .hero-dot-text {
    font-size: 9px;
  } */
}
</style>