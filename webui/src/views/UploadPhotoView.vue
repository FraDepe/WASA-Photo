<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			file: null,
			loggedId: null
		}
	},
	methods: {
		async uploadPhoto() {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.post("/users/" + localStorage.userid + "/photos", this.imageFile , {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.$router.push("/home");
			} catch (e) {
				this.$router.push("/error");
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		uploadedImage() {
			this.imageFile = this.$refs.image.files[0]
			this.file = URL.createObjectURL(this.$refs.image.files[0])

		},
		
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Upload your photo</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="mb-3">
			<label for="description" class="form-label">Picture</label><br />
			<input type="file" class="uploading-image" accept="image/png, image/jpg, image/jpeg" ref="image" @change="uploadedImage" >
		</div>
		
		<img v-if="this.file != null" :src="this.file" style="height: 400px;"/><br />

		<br />
		<div>
			<button v-if="!loading" :disabled="this.file == null" type="button" class="btn btn-primary" @click="uploadPhoto">
				Upload
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
