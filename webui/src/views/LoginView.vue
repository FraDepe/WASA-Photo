<script>
export default {
	data: function () {
		return {
			errormsg: null,
			username: "",
            user: {}
		}
	},
	methods: {
		doLogin: async function () {
            localStorage.clear          // perform a clear to logout and login
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session", this.username 
				);
				this.$router.push("/home");
                this.user = response.data;
                localStorage.userid = response.data["id"]
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="mb-3">
			<label for="description" class="form-label">Username</label>
			<input type="string" class="form-control" id="username" v-model="username" placeholder="Your username">
		</div>

		<div>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="doLogin">
				Login
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
