<script>
    export let url;
</script>

<div class="audio">
    <audio class="display-audio" controls preload="auto">
        <source src={url} />
    </audio>
</div>

<style lang="scss">
    .audio {
        margin: 0 -1rem 1rem;
        audio {
            border-radius: 0.3rem;
            width: 100%;
        }
    }
</style>
