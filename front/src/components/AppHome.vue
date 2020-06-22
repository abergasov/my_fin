<template>
  <div>
    <router-link to="/">Home</router-link> |
    <router-link to="/medium">Medium</router-link>
    <router-view />
  </div>
</template>

<script>
export default {
    name: 'AppHome',
    created() {
      this.askBackend('asd');
    }
};
</script>

<style scoped>
    h1  {
        color: green;
    }
</style>