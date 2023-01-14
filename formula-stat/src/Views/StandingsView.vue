<template>
    <div class="mt-5">
        <h1 class="is-size-3 mb-4 has-text-weight-semibold has-text-centered">Drivers Championship</h1>
        <div class="columns is-mobile is-vcentered box is-gapless">
            <div class="column is-4">
                <button @click="yearBackwards" class="button is-large is-primary is-pulled-left is-light is-rounded">&lt;</button>
            </div>
            <div class="column is-4">
                <h1 class="is-size-2 has-text-centered">{{ year }}</h1>
            </div>
                <div class="column is-4">
                    <button @click="yearForwards" class="button is-large is-pulled-right is-primary is-light is-rounded">&gt;</button>
                </div>
        </div>
        <button @click="showHideRounds" class="button is-small is-light">{{ showRounds ? "Hide" : "Show"}} Rounds</button>
        <div v-if="this.showRounds" class="columns is-mobile is-vcentered box is-gapless">
            <div class="column is-4">
                <button @click="roundBackwards" class="button is-large is-primary is-pulled-left is-light is-rounded">&lt;</button>
            </div>
            <div class="column is-4">
                <h1 class="is-size-2 has-text-centered">R. {{ round }}</h1>
            </div>
                <div class="column is-4">
                    <button @click="roundForwards" class="button is-large is-pulled-right is-primary is-light is-rounded">&gt;</button>
                </div>
        </div>
        <div class="mt-5 columns is-gapless is-mobile box is-vcentered">
            <div class="column is-3">
                <h1 class="is-size-6 has-text-grey has-text-weight-semibold has-text-centered">Position</h1>
            </div>
            <div class="column is-3">
                <h1 class="is-size-6 has-text-grey has-text-weight-semibold has-text-centered">Driver</h1>
            </div>
            <div class="column is-3">
                <h1 class="is-size-6 has-text-grey has-text-weight-semibold has-text-centered">Team</h1>
            </div>
            <div class="column is-3">
                <h1 class="is-size-6 has-text-grey has-text-weight-semibold has-text-centered">Points</h1>
            </div>
        </div>
        <div class="columns is-multiline is-gapless">
            <div v-for="driver in drivers" v-bind:key="driver.position" class="column is-full">
                <StandingsItemVue :driver="driver" />
            </div>
        </div>
    </div>
</template>

<script>
import StandingsItemVue from '@/components/StandingsItem.vue';
import Driver from '@/js/Driver'

export default {
    name: 'StandingsVue',
    components: {
        StandingsItemVue
    },
    methods: {
        yearBackwards() {
            if(this.year > 1950) {
                this.year-=1
                this.drivers = Driver.retrieveDrivers(this.year, 1)
                this.round=this.drivers[0].totalrounds
                this.drivers = Driver.retrieveDrivers(this.year, this.round)
            }
        },
        yearForwards() {
            if(this.year < 2023) {
                this.year+=1
                this.drivers = Driver.retrieveDrivers(this.year, 1)
                this.round=this.drivers[0].totalrounds
                this.drivers = Driver.retrieveDrivers(this.year, this.round)
            }
        },
        roundBackwards() {
            if(this.round > 1) {
                this.round-=1
                this.drivers = Driver.retrieveDrivers(this.year, this.round)
            }
        },
        roundForwards() {
            if(this.round < this.drivers[0].totalrounds) {
                this.round+=1
                this.drivers = Driver.retrieveDrivers(this.year, this.round)
            }
        },
        showHideRounds() {
            this.showRounds = !this.showRounds
        }

    },
    data() {

        var generatedDrivers = Driver.retrieveDrivers(2022, 22)

        return {
            year: 2022,
            round: 22,
            drivers: generatedDrivers,
            showRounds: false,
        }
    },
    mounted() {
        // console.log("MOUNTED")
        // this.drivers = Driver.retrieveDrivers(2022, 22)
    }
}
</script>