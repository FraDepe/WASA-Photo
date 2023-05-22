<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			photos: [],
			likes: [],
			comment: "",
			username: "",
		}
	},
	methods: {
		load() {
			return load
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users", {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.photos = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async uploadPhoto() {
			this.$router.push("/new");
		},
		async changeUser() {
			this.$router.push("/");
		},
		async deleteFountain(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/fountains/" + id);

				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
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
		},
		liked(photoid, p_userid) {
			this.errormsg = null;
			try {
				let response = this.$axios.get("/photos/" + photoid + "/likes/", {
					headers: {
						Authorization: p_userid
					}
				});
				likes = response.data
				for (let i = 0; i < likes.length; i++) {
					if (likes[i]["UserId"] == localStorage.userid) {
						return true;
					}
				}
				return false;
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Following stream</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="input-group me-2">
					<input type="string" class="form-control" v-model="username" placeholder="Search here for a user">
					<button type="button" class="btn btn-primary" @click="getUserProfile(this.username)" >
						Search
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">
						Upload a photo
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="changeUser">
						Change user
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="photos.length === 0">
			<div class="card-body">
				<p>No one uploaded a photo</p>
			</div>
		</div>

		<div class="card" >
			<div class="card-header">
				Photos
			</div>
			<div class="card-body" >
				<div class="container" ></div>
					<div class="row" >
						<div class="col-md-6" v-if="!loading" v-for="p in photos">
							<p class="card-text">
								<img :src="'data:image;base64,' + p.Picture" style="height: 250px;"/><br />
								User: {{ p.User_id }}<br />
								Date: {{ p.Date_time }}<br />
								Likes: {{ p.Likes }}<br />
								Comments: {{ p.Comments }}<br />
							<button type="button" class="btn btn-primary" @click="likePhoto(p.ID)">Like</button>
							</p>
							<div class="input-group mb-3">
								<input type="string" class="form-control" v-model="comment" placeholder="Write your comment here">
								<button type="button" class="btn btn-primary" @click="commentPhoto(p.ID)" >Post</button>
							</div>
						</div>
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
