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
            //box-shadow: 0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23);
        }
    }
</style>
