import { fail, type Actions } from '@sveltejs/kit';
import type {  PageServerLoad } from './$types';

type User={
    name:string
    email:string
    gender:string
}

/*export const load = (async ({ fetch }) => {
    const response=await fetch('http://localhost:8080/').then(r => r.json())
    return{response}
}) satisfies PageServerLoad;
*/

export const actions = {
	login: async ({  request }) => {
		const data = await request.formData();
		const email = data.get('email');
		const password = data.get('password');

        if (!email) {
            console.log(password)
			return fail(400, { email, missing: true });
		}
		
		return { success: true };

	}
	
} satisfies Actions;