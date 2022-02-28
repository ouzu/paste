<script>
    export let url;
    export let name;
</script>

<div class="image">
    <img src={url} alt={name} />
</div>

<style lang="scss">
    .image {
        margin: 0 -1rem 1rem;
        text-align: center;

        img {
            max-width: 100%;
            border-radius: 0.3rem;
            //box-shadow: 0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23);
        }
    }
</style>