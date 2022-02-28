<script>
    export let url;
</script>

<div class="audio">
    <audio class="display-audio" controls preload="auto">
        <source src={url} />
        <track kind="captions">
    </audio>
</div>

<style lang="scss">
    .audio {
        margin: 0 -1rem 1rem;
        audio {
            border-radius: 0.3rem;
            width: 100%;
            //box-shadow: 0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23);
        }
    }
</style>
