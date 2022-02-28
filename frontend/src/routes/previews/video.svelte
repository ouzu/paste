<script>
    export let url;
</script>

<div class="video">
    <video class="display-video" controls autoplay>
        <source src={url} />
        <track kind="captions" />
    </video>
</div>

<style lang="scss">
    .video {
        margin: 0 -1rem 1rem;
        video {
            border-radius: 0.3rem;
            width: 100%;
        }
    }
</style>
