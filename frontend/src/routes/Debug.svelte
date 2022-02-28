<script lang="ts">
    import { onMount } from 'svelte';
    import {generateKey, importKey, exportKey, generateIV, importIV, exportIV} from "../crypto/crypto";
    
    let key: CryptoKey;
    let keyString: string;
    let iv: Uint8Array;
    let ivString: string;

    async function init() {
        key = await generateKey();
        keyString = await exportKey(key);
        iv = await generateIV();
        ivString = await exportIV(iv);

        console.log()
    }

    onMount(init);

</script>

<div class="debug">
    Key: {keyString}<br>
    IV: {ivString}
</div>