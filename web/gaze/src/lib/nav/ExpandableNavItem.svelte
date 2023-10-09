<script>
    export let itemName = ""
    export let expand = false
</script>

<div class="p-5  cursor-pointer flex gap-x-4 text-slate-200 text-lg font-mono tracking-widest select-none " on:click={() => expand=!expand}>
    <div class="w-8">
        <div class="w-0 h-0 relative top-2
        border-l-[8px] border-l-transparent
        border-t-[8px] border-t-slate-400
        border-r-[8px] border-r-transparent
        transition duration-150 ease-in-out
        {expand === false ? '-rotate-90' : ''}
    ">
        </div>
    </div>


    <div>{itemName}</div>
</div>
