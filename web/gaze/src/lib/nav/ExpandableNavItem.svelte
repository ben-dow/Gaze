<script>
    export let itemName = ""
    export let expand = false
</script>

<div class="p-5 cursor-pointer flex gap-x-4 text-slate-300 text-lg font-mono tracking-widest select-none hover:bg-slate-800 first:hover:border-t-slate-200" >
    <div class="w-8 h-8 flex group/arrow" on:click={() => expand=!expand}>
        <div class="m-auto group-hover/arrow:border-t-slate-200
        border-l-[6px] border-l-transparent
        border-t-[6px] border-t-slate-400
        border-r-[6px] border-r-transparent
        transition duration-150 ease-in-out
        {expand === false ? '-rotate-90' : ''}
    ">
        </div>
    </div>


    <div class="h-8 my-auto hover:text-slate-50">{itemName}</div>
</div>
