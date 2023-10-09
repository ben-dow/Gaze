<script>
    export let itemName = ""
    export let expand = false
</script>

<div class="p-5  cursor-pointer flex gap-x-4 text-slate-200 text-lg font-mono tracking-widest select-none " on:click={() => expand=!expand}>
    <div class="w-8"></div>
    <div>{itemName}</div>
</div>
