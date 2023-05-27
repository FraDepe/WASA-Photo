<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			photos: [],
			likes: [],
			comment: "",
			usernameToSearch: "",
			following: false,
			photo: {
				id: 0,
				user_id: 0,
				picture: null,
				likes: 0,
				date_time: "",
				comments: 0
			},
			loggedId: 0
            
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photos/"+ this.$route.params.photoid, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.photo = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.loggedId = localStorage.userid
			this.loading = false;
		},
		async likePhoto(photoid) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.post("/photos/" + photoid + "/likes/" + localStorage.userid , null, {
					headers: {
						Authorization: localStorage.userid
					}
				});

				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async unlikePhoto(photoid) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/photos/" + photoid + "/likes/" + localStorage.userid ,null , {
					headers: {
						Authorization: localStorage.userid
					}
				});

				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async commentPhoto(photoid) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.post("/photos/" + photoid + "/comments/", this.comment , {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.comment = ""
				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getUserProfile(string) {
			localStorage.usernameToSearch = string
			this.$router.push("/user/"+ string);
			this.usernameToSearch = ""
			this.refresh()
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="input-group me-2">
					<input type="string" class="form-control" v-model="usernameToSearch" placeholder="Search here for a user">
					<button type="button" class="btn btn-primary" @click="getUserProfile(this.usernameToSearch)" >
						Search
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card">
			<div class="card-header">
				Photo dated {{ this.photo.date_time }} <br/> 
			</div>
			<div class="card-body">
                <p class="card-text">
                    <img :src="'data:image;base64,' + this.photo.picture" style="height: 250px;"/><br />
                    Likes: {{ this.photo.likes }} <br />
                    <div v-if="this.loggedId != this.photo.user_id">
                        <button type="button" class="btn btn-primary" @click="likePhoto()">Like</button>
                        <button type="button" class="btn btn-danger" @click="unlikePhoto()">Unlike</button>
                    </div>
                </p>
                <div class="input-group mb-3">
                    <input type="string" class="form-control" v-model="comment" placeholder="Write your comment here">
                    <button type="button" class="btn btn-primary" @click="commentPhoto()" >Post</button>
                </div>
			</div>
		</div>

		
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
