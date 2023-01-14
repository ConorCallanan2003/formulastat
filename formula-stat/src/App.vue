<template>
  <div class="m-2">
    <!-- <HomeVue></HomeVue> -->
    <StandingsView />
  </div>
</template>

<script>
// import HomeVue from './Views/Home.vue';
import StandingsView from './Views/StandingsView.vue';

export default {
  name: 'App',
  data () {
  },
  components: {
    // HomeVue
    StandingsView
}
}
</script>
