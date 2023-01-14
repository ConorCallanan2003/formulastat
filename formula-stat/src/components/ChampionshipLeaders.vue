<template>
    <div class="has-background-light box block">
        <div class="columns is-mobile is-multiline is-vcentered">
            <div class="column is-full">
                <h1 class="has-text-centered title has-text-grey-darker">{{ this.championship }}</h1>
            </div>
            <div class="column is-3 block " >
                <ChampionshipLeaderItemVue :name="this.names[1]" :imageSrc="this.imageUrls[1]" position="3" />
            </div>
            <div class="column is-6 block">
                <ChampionshipLeaderItemVue :name="this.names[0]" :imageSrc="this.imageUrls[0]" position="1" />
            </div>
            <div class="column is-3 block ">
                <ChampionshipLeaderItemVue :name="this.names[2]" :imageSrc="this.imageUrls[2]" position="2"/>
            </div>
        </div>
    </div>
</template>

<script>
import ChampionshipLeaderItemVue from './ChampionshipLeaderItem.vue';

export default {
    name: "ChampionshipLeaders",
    props: {
        championship: String,
        names: Array,
        imageUrls: Array
    },
    components: { ChampionshipLeaderItemVue },
}
</script>
