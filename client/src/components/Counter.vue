<template>
  <div class="counter-container">
    <h2>Interactive Vue Counter</h2>
    <p>Current count: <span class="count">{{ count }}</span></p>
    <div class="buttons">
      <button @click="increment" class="btn btn-primary">+</button>
      <button @click="decrement" class="btn btn-secondary">-</button>
      <button @click="reset" class="btn btn-danger">Reset</button>
    </div>
    <p class="status">{{ statusMessage }}</p>
  </div>
</template>

<script>
export default {
  name: 'Counter',
  data() {
    return {
      count: 0
    }
  },
  computed: {
    statusMessage() {
      if (this.count === 0) return 'Starting fresh!'
      if (this.count > 0) return `You're ${this.count} step(s) ahead!`
      return `You're ${Math.abs(this.count)} step(s) behind!`
    }
  },
  methods: {
    increment() {
      this.count++
    },
    decrement() {
      this.count--
    },
    reset() {
      this.count = 0
    }
  }
}
</script>

<style scoped>
.counter-container {
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  text-align: center;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.count {
  font-size: 2rem;
  font-weight: bold;
  color: #ffd700;
}

.buttons {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  margin: 1rem 0;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  font-weight: bold;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.btn-primary {
  background: #48bb78;
  color: white;
}

.btn-secondary {
  background: #ed8936;
  color: white;
}

.btn-danger {
  background: #f56565;
  color: white;
}

.status {
  margin-top: 1rem;
  font-style: italic;
  opacity: 0.9;
}
</style>