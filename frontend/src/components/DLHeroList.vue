<template>

    <div class="hero-list-wrapper">
        <div class="search-container">
            <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Search heroes..."
                class="search-input"
            />
            <button class="add-combo-btn" @click="addCombo" title="Add a Combo">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
            </button>
        </div>
        
        <div class="checkbox-container">
            <label 
                v-for="hero in sortedAndFilteredHeroes" 
                :key="hero.Heroes[0].ID" 
                class="checkbox-item" 
                :class="{'checkbox-item-checked': checkedHeroes.includes(hero.Heroes[0].ID)}" 
                :style="`--hero-color: ${hero.Heroes[0].Color}`"
            >
                <input 
                    type="checkbox" 
                    :value="hero.Heroes[0].ID"
                    v-model="checkedHeroes"
                />
                <img 
                    class="image"
                    :src="hero.Heroes[0].Image"
                />
                <span class="name">{{ hero.Heroes[0].Name }}</span>
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
  computed: {
    sortedAndFilteredHeroes() {
      if (!this.data) return [];
      
      // Filter by search query
      let filtered = this.data.filter(hero => {
        const name = hero.Heroes[0].Name.toLowerCase();
        const query = this.searchQuery.toLowerCase();
        return name.includes(query);
      });
      
      // Sort: checked heroes first, then alphabetically
      return filtered.sort((a, b) => {
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
    addCombo() {
      alert('Add a Combo');
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
    border: solid 2px transparent;
    color: lightslategray;
    border-radius: 5px;
    margin-bottom: 0.25rem;
}

.checkbox-item-checked {
    background-color: var(--hero-color);
    opacity: 0.75;
    color: white;
}

.checkbox-item-checked:hover {
    opacity: 1;
}

.checkbox-item:hover {
    border: solid 2px var(--hero-color);
    color: white;
}

.image {
  display: block;
  width: 20px;
  height: 20px;
  object-fit: contain;
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