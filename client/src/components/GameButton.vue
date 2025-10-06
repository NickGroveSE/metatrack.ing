<template>
  <a class="button-container" :href="path">
    <div class="button-image" :style="{ backgroundImage: `url(${logo})`, backgroundSize: 'cover',  backgroundOrigin: 'content-box', backgroundRepeat: 'no-repeat', backgroundPosition: 'center' }"></div>
    <div class="button-title">{{ name }}</div>
    <div class="button-title-2">{{ namesl }}</div>
    <div class="coming-soon-layer" v-show="!active">Coming Soon</div>
  </a>
</template>

<script>
export default {
  props: {
    logo: String,
    path: String,
    name: String,
    namesl: String,
    active: Boolean,
  }
}
</script>

<style scoped>
.button-container {
  flex-basis: 50%;
  position: relative;
  text-decoration: none;
}

.button-image {
  aspect-ratio: 1;
  border: 5px solid #fbbf24;
  border-radius: 25px;
  padding: 20px;
}

.button-title {
  margin-top: 10px;
  padding: 0 10px;
  color: #fbbf24;
  font-weight: 700;
  font-style: italic;
  font-size: 30px;
  text-align: center;
}

.button-title-2 {
  padding: 0 10px;
  color: #fbbf24;
  font-weight: 700;
  font-style: italic;
  font-size: 30px;
  text-align: center;
}

.coming-soon-layer {
  display: flex;
  position: absolute;
  top: 0;
  left: 0;
  border: 5px solid transparent;
  border-radius: 25px;
  background-color: rgba(0, 0, 0, 0.6);
  width: 100%;
  aspect-ratio: 1;
  color: #fbbf24;
  font-weight: 700;
  font-style: italic;
  font-size: 36px;
  text-align: center;
  align-items: center;
  justify-content: center;
}
</style>