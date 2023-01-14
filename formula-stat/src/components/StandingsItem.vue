<template>
    <div class="columns is-gapless is-mobile box is-vcentered">
        <div class="column is-1">
            <h1 class="is-size-5 has-text-weight-semibold">{{this.driver.position}}</h1>
        </div>
        <div class="column is-5">
            <h1 class="is-size-5 has-text-grey has-text-weight-semibold">{{this.driver.lastName}}</h1>
        </div>
        <div class="column is-5">
            <h1 class="is-size-6 has-text-grey has-text-weight-semibold">{{this.driver.team}}</h1>
        </div>
        <div class="column is-1">
            <h1 class="is-size-5 has-text-weight-semibold">{{this.driver.points}}</h1>
        </div>
    </div>
</template>

<script>
export default {
    name: 'StandingsItem',
    props: {
        driver: Object
    }
}
</script>