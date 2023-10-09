<script>
    import testImg from "$lib/test_pic.png"
    import { Icon, Heart,ArrowDownCircle,Share, EllipsisHorizontal } from "svelte-hero-icons";
</script>

<div class="flex flex-col bg-gray-800 w-64 h-64 shadow-black shadow-md select-none border-slate-600 border-solid border-4 rounded-2xl border-double hover:border-solid cursor-pointer hover:bg-opacity-75  text-slate-300 font-mono justify-evenly">
    <div class="text-sm w-full text-center">
        <div>test_pic.png</div>
    </div>

    <div class="h-36 w-full object-center">
        <img class="max-h-36 object-center mx-auto" src={testImg} alt="test" \>
    </div>

    <div class="flex flex-col justify-center px-4">
        <div class="flex flex-row w-full justify-between text-2xs">
            <div class="h-full hover:bg-gray-700 p-4 rounded">
                <Icon src="{Heart}" class="max-h-4 hover:bg-gray-700 "/>
            </div>
            <div class="h-full hover:bg-gray-700 p-4 rounded">
                <Icon src="{Share}" class="max-h-4"/>
            </div>
            <div class="h-full hover:bg-gray-700 p-4 rounded">
                <Icon src="{ArrowDownCircle}" class="max-h-4"/>
            </div>
            <div class="h-full hover:bg-gray-700 p-4 rounded">
                <Icon src="{EllipsisHorizontal}" class="max-h-4"/>
            </div>
        </div>
    </div>
</div>