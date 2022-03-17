<script lang="ts">
	import "normalize.css/normalize.css";
	import DropZone from "svelte-atoms/DropZone.svelte";

	import ProgressBar from "svelte-progress-bar";

	import {
		generateKey,
		exportKey,
		generateIV,
		exportIV,
		encrypt,
		encryptString,
	} from "../crypto/crypto";

	import { _ } from "../i18n";

	let files = [];

	let progress;

	let uploading = false;

	document.onpaste = function (event) {
		const items = event.clipboardData.items;
		for (var i = 0; i < items.length; i++) {
			const item = items[i];

			if (item.kind === "file") {
				fileHandler(item.getAsFile());
			} else if (item.kind === "string" && item.type === "text/plain") {
				items[i].getAsString((s) => {
					const encoder = new TextEncoder();
					let f = new File([encoder.encode(s)], "paste.txt", {
						type: "text/plain",
					});

					fileHandler(f);
				});
			} else if (item.kind === "string" && item.type === "text/html") {
				items[i].getAsString((s) => {
					const encoder = new TextEncoder();
					let f = new File([encoder.encode(s)], "paste.html", {
						type: "text/html",
					});
					fileHandler(f);
				});
			}
		}
	};

	async function fileHandler(file) {
		uploading = true;

		const fileName = file.name;
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
			requestData.append("name", await encryptString(fileName, iv, key));
			requestData.append("type", await encryptString(fileType, iv, key));

			let request = new XMLHttpRequest();
			request.open("PUT", "/api/files/" + ivStr);

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
				uploading = false;
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
		<DropZone
			title={$_("home.drag")}
			buttonTitle={$_("home.select")}
			activeTitle={$_("home.drop")}
			on:drop={dropHandler}
			on:change={dropHandler}
			disabled={uploading}
		/>
	</div>
	{#if files.length > 0}
		<h3>{$_("home.files")}</h3>
		{#each files as { name, url, error }}
			{#if error === ""}
				<p><a href={url} target="_blank">{name}</a></p>
			{:else}
				<p>{name} - Error: {error}</p>
			{/if}
		{/each}
	{/if}
</div>

<style lang="scss">
	.container {
		background-color: var(--ac-1);
		border-radius: 0.3rem;
		padding: 1rem;
		margin: 0 -1rem;
	}
</style>
