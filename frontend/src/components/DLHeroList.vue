<template>
    <div class="hero-list-wrapper">

        <div class="search-container">
            <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Search heroes..."
                class="search-input"
            />
            <button :class="`hero-list-btn ${filtersModalOpen ? 'open' : ''}`"  @click="openModal">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"></polygon>
                </svg>
            </button>
            <div v-if="filtersModalOpen" class="filters-container">
              <div class="filter-group">
                <label class="filter-label">Timeframe</label>
                <select v-model="filters.timeframe" class="filter-select">
                  <option value="24h">Last 24 Hours</option>
                  <option value="48h">Last 48 Hours</option>
                  <option value="168h">Last 7 Days</option>
                  <option value="336h">Last 2 Weeks</option>
                </select>
              </div>

              <div class="filter-group">
                <label class="filter-label">Min Rank</label>
                <select v-model="filters.minBadge" class="filter-select">
                  <option value="0">Obscurus</option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(10+badgeNumber).toString()" 
                    :key="(10+badgeNumber).toString()">
                      {{ "Initiate " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(20+badgeNumber).toString()" 
                    :key="(20+badgeNumber).toString()">
                      {{ "Seeker " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(30+badgeNumber).toString()" 
                    :key="(30+badgeNumber).toString()">
                      {{ "Alchemist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(40+badgeNumber).toString()" 
                    :key="(40+badgeNumber).toString()">
                      {{ "Arcanist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(50+badgeNumber).toString()" 
                    :key="(50+badgeNumber).toString()">
                      {{ "Ritualist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(60+badgeNumber).toString()" 
                    :key="(60+badgeNumber).toString()">
                      {{ "Emissary " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(70+badgeNumber).toString()" 
                    :key="(70+badgeNumber).toString()">
                      {{ "Archon " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(80+badgeNumber).toString()" 
                    :key="(80+badgeNumber).toString()">
                      {{ "Oracle " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(90+badgeNumber).toString()" 
                    :key="(90+badgeNumber).toString()">
                      {{ "Phantom " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(100+badgeNumber).toString()" 
                    :key="(100+badgeNumber).toString()">
                      {{ "Ascendant " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(110+badgeNumber).toString()" 
                    :key="(110+badgeNumber).toString()">
                      {{ "Eternus " + badgeNumber.toString() }}
                  </option>
                </select>
              </div>
              
              <div class="filter-group">
                <label class="filter-label">Max Rank</label>
                <select v-model="filters.maxBadge" class="filter-select">
                  <option value="0">Obscurus</option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(10+badgeNumber).toString()" 
                    :key="(10+badgeNumber).toString()">
                      {{ "Initiate " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(20+badgeNumber).toString()" 
                    :key="(20+badgeNumber).toString()">
                      {{ "Seeker " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(30+badgeNumber).toString()" 
                    :key="(30+badgeNumber).toString()">
                      {{ "Alchemist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(40+badgeNumber).toString()" 
                    :key="(40+badgeNumber).toString()">
                      {{ "Arcanist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(50+badgeNumber).toString()" 
                    :key="(50+badgeNumber).toString()">
                      {{ "Ritualist " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(60+badgeNumber).toString()" 
                    :key="(60+badgeNumber).toString()">
                      {{ "Emissary " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(70+badgeNumber).toString()" 
                    :key="(70+badgeNumber).toString()">
                      {{ "Archon " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(80+badgeNumber).toString()" 
                    :key="(80+badgeNumber).toString()">
                      {{ "Oracle " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(90+badgeNumber).toString()" 
                    :key="(90+badgeNumber).toString()">
                      {{ "Phantom " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(100+badgeNumber).toString()" 
                    :key="(100+badgeNumber).toString()">
                      {{ "Ascendant " + badgeNumber.toString() }}
                  </option>
                  <option 
                    v-for="badgeNumber in badgeNumbering" 
                    :value="(110+badgeNumber).toString()" 
                    :key="(110+badgeNumber).toString()">
                      {{ "Eternus " + badgeNumber.toString() }}
                  </option>
                </select>
              </div>

              <button @click="applyFilters" class="apply-filters-btn">
                Apply Filters
              </button>
            </div>
            <button class="hero-list-btn" @click="dispatchCombo" title="Add a Combo">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
            </button>
        </div>

        <div :class="`select-all-button ${checkedHeroes.length == data.length ? 'selected' : ''}`" @click="dispatchSelectAll">Select All Heroes</div>
        
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
    },
    queryParams: Object
  },
  data() {
    return {
      data: this.heroData,
      checkedHeroes: [...this.setCheckedHeroes],
      searchQuery: '', 
      filtersModalOpen: false,
      filters: {
        timeframe: this.queryParams.timeframe,
        minBadge: this.queryParams.minBadge,
        maxBadge: this.queryParams.maxBadge
      },
      badgeNumbering: [1,2,3,4,5,6]
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
        
        // Both combos - sort by combo number
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
      alert("Add Hero Combo or Matchup In Development")
      // const event = new CustomEvent('combo-added', {
      //   detail: {
      //     checkedHeroes: this.checkedHeroes,
      //     combo: "2-11"
      //   },
      //   bubbles: true
      // });
      // this.$el.dispatchEvent(event);
    },
    dispatchSelectAll() {
      this.checkedHeroes = this.data.map(hero => hero.Heroes[0].ID)
    },
    applyFilters() {
      this.filtersModalOpen = false
      const event = new CustomEvent('apply-filters', {
        detail: {
          filters: this.filters,
          checkedHeroes: this.checkedHeroes
        },
        bubbles: true
      })
      this.$el.dispatchEvent(event)
    },
    openModal() {
      this.filtersModalOpen = !this.filtersModalOpen
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
    position: relative;
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
    flex-basis: 60%;
}

.hero-list-btn {
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

.select-all-button {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.25rem;
    border-radius: 8px;
    border: 1px solid rgba(31,41,55, 0.7);
    background-color: rgba(17,24,39, 0.5);
    color: lightslategray;
    cursor: pointer;
    transition: all 0.2s;
}

.selected {
    background-color: rgba(251, 191, 36, 0.2);
    border-color: rgba(251, 191, 36, 0.6);
    color: rgb(251, 191, 36);
    transform: scale(1.05);
} 

.hero-list-btn:hover {
    background-color: rgba(251, 191, 36, 0.2);
    border-color: rgba(251, 191, 36, 0.6);
    color: rgb(251, 191, 36);
    transform: scale(1.05);
}

.open {
    background-color: rgba(251, 191, 36, 0.2);
    border-color: rgba(251, 191, 36, 0.6);
    color: rgb(251, 191, 36);
    transform: scale(1.05);
}

.hero-list-btn:active {
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
    max-height: 420px;
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

.filters-container {
  background: #111827;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(31,41,55, 0.7);
  width: 300px;
  float: right;
  position: absolute;
  z-index: 1;
  top: 50px; 
}

/* Filter Group Styling */
.filter-group {
  margin-bottom: 10px;
}

.filter-label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: #e0e0e0;
  margin-bottom: 8px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.filter-select {
  width: 100%;
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.05);
  border: 2px solid rgba(249, 168, 38, 0.3);
  border-radius: 8px;
  color: #ffffff;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg width='12' height='8' viewBox='0 0 12 8' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1L6 6L11 1' stroke='%23f9a826' stroke-width='2' stroke-linecap='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 16px center;
  padding-right: 44px;
}

.filter-select:hover {
  border-color: rgba(249, 168, 38, 0.6);
  box-shadow: 0 0 15px rgba(249, 168, 38, 0.2);
}

.filter-select:focus {
  outline: none;
  border-color: #f9a826;
  box-shadow: 0 0 20px rgba(249, 168, 38, 0.3);
}

.filter-select optgroup {
  background: #1a1a2e;
  color: #06b6d4;
  padding: 10px;
}

.filter-select option {
  background: #1a1a2e;
  color: #ffffff;
  padding: 10px;
}

.apply-filters-btn {
  width: 100%;
  padding: 14px 24px;
  margin-top: 12px;
  background: linear-gradient(135deg, #f9a826 0%, #06b6d4 100%);
  border: none;
  border-radius: 8px;
  color: #ffffff;
  font-size: 16px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1px;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(249, 168, 38, 0.4);
}

.apply-filters-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 25px rgba(249, 168, 38, 0.6);
}

.apply-filters-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 10px rgba(249, 168, 38, 0.4);
}

</style>