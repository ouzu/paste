<script lang="ts">
	import { onMount } from "svelte";

	import "normalize.css/normalize.css";
	import DropZone from "svelte-atoms/DropZone.svelte";

	import ProgressBar from "svelte-progress-bar";

	import { textToArrayBuffer, arrayBufferToText } from "../crypto/helper";
	import {
		generateKey,
		importKey,
		exportKey,
		generateIV,
		importIV,
		exportIV,
		encrypt,
		decrypt,
		encryptString,
		decryptString,
	} from "../crypto/crypto";

	let files = [];

	let progress;

	document.onpaste = function (event) {
		const items = event.clipboardData.items;
		for (var i = 0; i < items.length; i++) {
			const item = items[i];
			console.log(item);

			if (item.kind === "file") {
				fileHandler(item.getAsFile());
			} else if (item.kind === "string" && item.type === "text/plain") {
				items[i].getAsString((s) => {
					console.log(s);
					const encoder = new TextEncoder();
					let f = new File([encoder.encode(s)], "paste.txt");

					fileHandler(f);
				});
			}
		}
	};

	async function fileHandler(file) {
		const fileName =  file.name;
		const fileType = file.type;

		let reader = new FileReader();
			let iv = generateIV();
			let ivStr = exportIV(iv);
			let key = await generateKey();
			let keyStr = await exportKey(key);

			reader.onload = async (e) => {
				let data = e.target.result;

				let requestData = new FormData();

				requestData.append(
					"file",
					new File([await encrypt(data, iv, key)], ivStr)
				);
				requestData.append(
					"name",
					await encryptString(fileName, iv, key)
				);
				requestData.append(
					"type",
					await encryptString(fileType, iv, key)
				);

				console.log("uploading " + fileName + " (" + fileType + ")");

				// fetch() doesn't support progress :(
				/*fetch("//localhost:8080/api/files/" + ivStr, {
					method: "PUT",
					body: requestData,
				}).then((r) => {
					if (r.ok) {
						files.push({
							name: fileName,
							url: "/#/files/" + ivStr + "/" + keyStr,
							error: "",
						});
					} else {
						console.log(r);
						files.push({
							name: fileName,
							url: "",
							error: r.statusText,
						});
					}
					files = files;
				});*/

				let request = new XMLHttpRequest();
				request.open("PUT", "//localhost:8080/api/files/" + ivStr);

				request.upload.addEventListener("progress", function (e) {
					progress.setWidthRatio(e.loaded / e.total);
				});

				request.addEventListener("load", function (e) {
					if (request.status == 200) {
						files.push({
							name: fileName,
							url: "/#/files/" + ivStr + "/" + keyStr,
							error: "",
						});
					} else {
						console.log(request.responseText);
						files.push({
							name: fileName,
							url: "",
							error: "could not upload",
						});
					}
					progress.setWidthRatio(0);
					files = files;
				});

				request.send(requestData);
			};

			reader.readAsArrayBuffer(file);
	}

	async function dropHandler(e) {
		const file = e.dataTransfer
			? e.dataTransfer.files[0]
			: e.target.files[0];

		if (file) {
			fileHandler(file);
		} else {
			console.error("Event contains no file " + file);
			console.debug(e);
		}
	}
</script>

<div class="home">
	<ProgressBar bind:this={progress} color="#abc4ff" />

	<div class="container">
		<DropZone on:drop={dropHandler} on:change={dropHandler} />
	</div>
	{#if files.length > 0}
		<h3>Files:</h3>
		{#each files as { name, url, error }}
			{#if error === ""}
				<p><a href={url} download={name}>{name}</a></p>
			{:else}
				<p>{name} - Error: {error}</p>
			{/if}
		{/each}
	{/if}
</div>

<style lang="scss">
	$color1: #edf2fb;
	$color2: #d7e3fc;
	$color3: #abc4ff;

	.container {
		background-color: $color1;
		border-radius: 0.3rem;
		padding: 1rem;
		margin: 0 -1rem;
		//box-shadow: 0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23);
	}
</style>