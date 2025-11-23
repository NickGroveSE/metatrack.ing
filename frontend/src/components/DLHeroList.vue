<template>

    <div class="checkbox-container">
        <label 
            v-for="(hero, i) in data" 
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

</template>

<script>

export default {
  name: 'DeadlockHeroList',
  props: {
    heroData: Object,
  },
  data() {
    return {
      data: this.heroData,
      checkedHeroes: []
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
  }
};

</script>

<style scoped>
.checkbox-container {
    width: 300px;
    max-height: 500px;
    overflow-y: auto;
    padding: 0.5rem;
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