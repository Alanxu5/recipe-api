<script>
export default {
  name: 'App',
}
</script>

<template>
  <div id="app">
    <div class="header">
      <div class="icon">
        <RouterLink to="/">
          <a class="router">
            Icon
          </a>
        </RouterLink>
      </div>
      <div class="link">
        <RouterLink to="/">
          <a class="router">
            Home
          </a>
        </RouterLink>
      </div>
      <div class="link">
        <RouterLink to="/recipes">
          <a class="router">
            Recipes
          </a>
        </RouterLink>
      </div>
      <div />
      <div class="link">
        <RouterLink 
          to="submit" 
          append>
          <a class="router">
            Submit
          </a>
        </RouterLink>
      </div>
    </div>
    <RouterView />
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.header {
  display: grid;
  grid-template-columns: minmax(2rem, 5%) minmax(9rem, 15%) minmax(9rem, 15%) 1fr minmax(9rem, 15%);
  background-color: lightblue;
  height: 3rem;
}

.icon {
  background-color: orange;
}

.link {
  display: grid;
  justify-items: center;
  align-items: center;
}

.link:hover {
  cursor: pointer;
  background-color: gray;
}

a:-webkit-any-link {
  color: black;
  text-decoration-line: none;
}
</style>
