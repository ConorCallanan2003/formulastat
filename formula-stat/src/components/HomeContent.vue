<template>
  <div class="mb-5 mt-4">
    <h1 class="title has-text-centered is-secondary has-text-weight-bold has-text-dark">Formula 1 Stats</h1>
  </div>
  <div class="box has-background-danger">
    <div class="columns is-mobile is-gapless is-multiline">
      <div class="column is-full">
        <h1 class="subtitle has-text-weight-semibold has-text-white-ter">Next Race</h1>
      </div>
      <div class="column is-full">
        <h1 class="subtitle has-text-weight-semibold has-text-white-ter is-size-3">Sao Paolo, Brazil</h1>
      </div>
    </div>
  </div>
  <ChampionshipLeadersVue :names=this.names :imageUrls=this.images championship="Drivers Championship"/>
  <StandingsVue title="Drivers Standings" backgroundColor="box has-background-primary" :listItems=this.drivers />
  <StandingsVue title="Constructors Standings" backgroundColor="box has-background-info" :listItems=this.teams />
</template>

<script>
import StandingsVue from './Standings.vue';
import ChampionshipLeadersVue from './ChampionshipLeaders.vue';
export default {
  name: 'HomeContent',
  data () {
    return {
      drivers: ["Hamilton", "Russell", "Norris", "Piastri", "Leclerc", "Sainz", "Stroll", "Alonso", "Tsunoda", "De Vrij", "Ocon", "Gasly", "Zhou", "Bottas", "Hulkenburg", "Magnussen", "Verstappen", "Perez", "Albon", "Sargeant"],
      teams: ["Mercedes", "McLaren", "Ferrari", "Aston Martin", "Alpine", "Alpha Tauri", "Alfo Romeo", "Haas", "Red Bull", "Williams"],
      images: ["https://cdn.racingnews365.com/Riders/Hamilton/_570x570_crop_center-center_none/f1_2022_lh_mer_lg.png?v=1647864750", "https://e2.365dm.com/f1/drivers/256x256/h_full_1496.png", "https://www.f1fantasytracker.com/Images/Drivers/2021/Headshots/Norris.png"],
      names: ["Hamilton", "Russell", "Norris"]
    }
  },
  props: {
    msg: String
  },
  components: {
    StandingsVue,
    ChampionshipLeadersVue
  }
}
</script>
