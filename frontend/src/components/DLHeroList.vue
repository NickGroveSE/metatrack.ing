<template>

    <div class="hero-list-wrapper">
        <div class="search-container">
            <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Search heroes..."
                class="search-input"
            />
            <button class="add-combo-btn" @click="dispatchCombo" title="Confirm Combo">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
            </button>
        </div>
        
        <div class="checkbox-container">
            <label 
                v-for="item in sortedAndFilteredHeroes" 
                :key="getItemKey(item)" 
                class="checkbox-item" 
                :class="getCheckboxClass(item)" 
                :style="getItemStyle(item)"
            >
                <input 
                    type="checkbox" 
                    v-if="item.Heroes.length === 1"
                    :value="item.Heroes[0].ID"
                    v-model="checkedHeroes"
                />
                <input 
                    type="checkbox"
                    v-else
                    :checked="true"
                    :disabled="true"
                />
                <img 
                    v-if="item.Heroes.length === 1"
                    class="image"
                    :src="item.Heroes[0].Image"
                />
                <template v-else>
                    <img 
                        v-for="hero in item.Heroes"
                        :key="hero.ID"
                        class="image combo-image"
                        :src="hero.Image"
                    />
                </template>
                <span class="name">{{ item.Heroes.length === 1 ? item.Heroes[0].Name : `Combo ${getComboNumber(item)}` }}</span>
            </label>
        </div>
    </div>

</template>

<script>

export default {
  name: 'DeadlockHeroList',
  props: {
    heroData: Object,
    setCheckedHeroes: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      data: this.heroData,
      checkedHeroes: [...this.setCheckedHeroes],
      searchQuery: ''
    };
  },
  mounted() {
    this.$el.addEventListener('hero-list-updated', (event) => {
      // This updates internal data, not the prop
      this.data = event.detail;
    });
  },
  computed: {
    sortedAndFilteredHeroes() {
      if (!this.data) return [];
      
      // Filter by search query
      let filtered = this.data.filter(hero => {
        const name = hero.Heroes[0].Name.toLowerCase();
        const query = this.searchQuery.toLowerCase();
        return name.includes(query);
      });
  
      // Sort: combos first, then checked heroes, then unchecked heroes
      return filtered.sort((a, b) => {
        const aIsCombo = a.Heroes.length > 1;
        const bIsCombo = b.Heroes.length > 1;
        
        // Combos always first
        if (aIsCombo && !bIsCombo) return -1;
        if (!aIsCombo && bIsCombo) return 1;
        
        // Both combos - sort alphabetically by combo number
        if (aIsCombo && bIsCombo) {
          return this.getComboNumber(a) - this.getComboNumber(b);
        }
      
        // Both are heroes - sort by checked status
        const aChecked = this.checkedHeroes.includes(a.Heroes[0].ID);
        const bChecked = this.checkedHeroes.includes(b.Heroes[0].ID);
        
        if (aChecked && !bChecked) return -1;
        if (!aChecked && bChecked) return 1;
        
        // If both checked or both unchecked, sort alphabetically
        return a.Heroes[0].Name.localeCompare(b.Heroes[0].Name);
      });
    }
  },
  methods: {
    dispatchCombo() {
      const event = new CustomEvent('combo-added', {
        detail: {
          checkedHeroes: this.checkedHeroes,
          combo: "2-11"
        },
        bubbles: true
      });
      this.$el.dispatchEvent(event);
    },
    getItemKey(item) {
      if (item.Heroes.length === 1) {
        return item.Heroes[0].Name
      }

      const names = item.Heroes.map(h => h.Name);

      return names.join("-")
    },
    getCheckboxClass(item) {
      if (item.Heroes.length === 1) {
          return {
              'checkbox-item-checked': this.checkedHeroes.includes(item.Heroes[0].ID)
          };
      }
      
      return {
          'checkbox-item-checked': true
      };
    },
    getItemStyle(item) {
      // Single hero - use solid color
      if (item.Heroes.length === 1) {
          return `--hero-color: ${item.Heroes[0].Color}`;
      }
      
      // Combo - create linear gradient with hard stops
      const colors = item.Heroes.map(h => h.Color);
      const stopPercentage = 100 / colors.length;
      
      let gradientStops = [];
      colors.forEach((color, index) => {
          const start = index * stopPercentage;
          const end = (index + 1) * stopPercentage;
          
          // Add both start and end points for hard transition
          gradientStops.push(`${color} ${start}%`);
          gradientStops.push(`${color} ${end}%`);
      });
      
      return `--hero-color: linear-gradient(135deg, ${gradientStops.join(', ')})`;
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
  },
  watch: {
    checkedHeroes: {
      handler(newVal) {
        console.log("Before Checked Hero Change Emission")
        // Dispatch a native DOM event that Astro can listen to
        const event = new CustomEvent('checked-heroes-changed', {
          detail: newVal,
          bubbles: true
        });
        this.$el.dispatchEvent(event);
      },
      deep: true
    },
    heroData: {
      handler(newVal) {
        this.data = newVal;
      },
      deep: true,
      immediate: true
    },
  }
};

</script>

<style scoped>
.hero-list-wrapper {
    width: 300px;
}

.search-container {
    margin-bottom: 0.75rem;
    display: flex;
    gap: 0.5rem;
}

.search-input {
    flex: 1;
    padding: 0.75rem;
    border-radius: 8px;
    border: 1px solid rgba(31,41,55, 0.7);
    background-color: rgba(17,24,39, 0.5);
    color: white;
    font-size: 0.95rem;
    transition: all 0.2s;
    box-sizing: border-box;
}

.add-combo-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.75rem;
    border-radius: 8px;
    border: 1px solid rgba(31,41,55, 0.7);
    background-color: rgba(17,24,39, 0.5);
    color: lightslategray;
    cursor: pointer;
    transition: all 0.2s;
    min-width: 48px;
}

.add-combo-btn:hover {
    background-color: rgba(251, 191, 36, 0.2);
    border-color: rgba(251, 191, 36, 0.6);
    color: rgb(251, 191, 36);
    transform: scale(1.05);
}

.add-combo-btn:active {
    transform: scale(0.95);
}

.search-input::placeholder {
    color: lightslategray;
}

.search-input:focus {
    outline: none;
    border-color: rgba(251, 191, 36, 0.6);
    box-shadow: 0 0 0 3px rgba(251, 191, 36, 0.1);
}

.checkbox-container {
    max-height: 450px;
    overflow-y: auto;
    padding: 1rem;
    border-radius: 12px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(31,41,55, 0.7);
}

.checkbox-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.5rem 0.75rem;
    cursor: pointer;
    transition: background 0.2s;
    color: lightslategray;
    border-radius: 5px;
    margin-bottom: 0.25rem;
}

.checkbox-item-checked {
    background: var(--hero-color);
    opacity: 0.75;
    color: white;
}

.checkbox-item-checked:hover {
    opacity: 1;
}

.checkbox-item:hover {
    box-shadow: 0 0 0 2px rgba(251, 191, 36, 0.6);
    color: white;
}

.image {
  display: block;
  width: 20px;
  height: 20px;
  object-fit: contain;
}

.combo-image {
    display: inline-block;
    margin-right: -8px;  /* Negative margin for overlap */
}

.combo-image:last-child {
    margin-right: 0;
}

.checkbox-item input {
    cursor: pointer;
}

.name {
    text-shadow: 
        -1px -1px 0 rgba(17,24,39, 128), 
        1px -1px 0 rgba(17,24,39, 128), 
        -1px 1px 0 rgba(17,24,39, 128), 
        1px 1px 0 rgba(17,24,39, 128);
}

</style>