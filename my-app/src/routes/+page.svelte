<script lang="ts">
    import Header from '$lib/components/Header.svelte';
import type { PageData, ActionData } from './$types';
    
    let {data,form}:{data:PageData,form:ActionData}=$props();
    let user=$state({email:'',password:''})
    

</script>
<div class="body_left">
	<div class="body_left_in">
		<Header/>
	</div>	
</div>
<div>

</div>
<style>
	.body_left{
		width: 80px;
		position: absolute;
		top: 0;
		left: 0;
		bottom: 0;
	}
</style>