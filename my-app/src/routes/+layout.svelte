<script lang="ts">
    
    let { data,children }= $props();
</script>
<div class="body">
    {@render children()}
</div>
<style>
    .body{
        background-color: #474747;
        border-radius: 8px;
        padding-left: 80px;
        position: relative;
        box-shadow: 0 2px 6px 0 rgba(0,0,0,.11),0 0 35px 0 rgba(0,0,0,.06);
        -webkit-box-shadow: 0 2px 6px 0 rgba(0,0,0,.11),0 0 35px 0 #3b3a3b;
    }
</style>