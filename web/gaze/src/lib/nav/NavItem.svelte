<script>
    export let itemName = ""
    export let expand = false
</script>

<div class="p-5  cursor-pointer flex gap-x-4 text-slate-300 text-lg font-mono tracking-widest select-none hover:text-slate-50 hover:bg-slate-800 rounded " on:click={() => expand=!expand}>
    <div class="w-8"></div>
    <div>{itemName}</div>
</div>
