<template>
    <div>
        <columns class="is-multiline is-mobile">
            <column class="is-full">
                <h1 class="title has-text-centered">{{ this.position }}</h1>
                <img :src=this.imageSrc />
                <h1 class="subtitle has-text-centered has-text-weight-semibold">{{ this.name }}</h1>
            </column>
            
        </columns>
    </div>
</template>


<script>
export default {
    name: 'ChampionshipLeaderItem',
    props: {
        position: String,
        imageSrc: String,
        name: String
    }
}
</script>