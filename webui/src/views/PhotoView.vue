<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			photos: [],
			listLikes: [],
			like: {
				photoid: 0,
				userid: 0
			},
			liked: null,
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
			comments: [],
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

			try {
				let response = await this.$axios.get("/photos/" + this.$route.params.photoid + "/comments/", {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.comments = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}

			try {
				let response = await this.$axios.get("/photos/" + this.$route.params.photoid + "/likes/" + localStorage.userid, {
					headers: {
						Authorization: this.photo.user_id
					}
				});
				this.like = response.data
			} catch (e) {
				this.errormsg = e.toString();
			}

			if (this.like.userid == 0) {
				this.liked = false
			} else {
				this.liked = true
			}

			try {
				let response = await this.$axios.get("/photos/" + this.$route.params.photoid + "/likes/", {
					headers: {
						Authorization: this.photo.user_id
					}
				});
				this.listLikes = response.data
			} catch (e) {
				this.errormsg = e.toString();
			}

            this.loggedId = localStorage.userid
			this.loading = false;
		},

		async deletePhoto() {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/photos/" + this.photo.id, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.$router.push("/home");
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
				await this.$axios.delete("/photos/" + photoid + "/likes/" + localStorage.userid , {
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
			this.refresh()
			this.loading = false;
		},

		async deleteComment(photoid, commentid) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/photos/" + photoid + "/comments/" + commentid, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.refresh()
			this.loading = false;
		},


	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div v-if="this.errormsg==null">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<br />

		<div class="card">
			<div class="card-header d-flex justify-content-between align-items-center">
				<h6 class="mb-0">Photo dated {{ this.photo.date_time }}</h6> 
				<button type="button" class="btn btn-danger" @click="deletePhoto()" v-if="this.photo.user_id == this.loggedId">Delete</button> <br/> 
			</div>
			<div class="card-body">

                <p class="card-text">
                    <img :src="'data:image;base64,' + this.photo.picture" style="height: 250px;"/><br />
                    Likes: {{ this.photo.likes }} <br />
                    <div v-if="this.loggedId != this.photo.user_id">
                        <button type="button" class="btn btn-primary" @click="likePhoto(this.photo.id)" v-if="liked == false">Like</button>
                        <button type="button" class="btn btn-danger" @click="unlikePhoto(this.photo.id)" v-else>Unlike</button>
                    </div>
                </p>

                <div class="input-group mb-3">
                    <input type="string" class="form-control" v-model="comment" placeholder="Write your comment here">
                    <button type="button" :disabled="this.comment == '' " class="btn btn-primary" @click="commentPhoto(this.photo.id)" >Post</button>
                </div>

				<div class="dropdown">
					<button type="button" class="btn btn-warning dropdown-toggle" data-bs-toggle="dropdown">
						Likes
					</button>
					<div class="dropdown-menu">
						<a class="dropdown-item disabled" v-for="l in this.listLikes"> {{ l.UserId }}</a>
					</div>
				</div>

				<br />

				<div class="card">
					<div class="box">
						<p class="card-text" v-for="c in comments">
							{{ c.UserId }} <br />
							{{ c.Text }} 
							<button type="button" class="btn btn-danger" @click="deleteComment(this.photo.id ,c.ID)" v-if="c.UserId == this.loggedId">Delete</button>
							<br />
							<hr>
						</p>
					</div>
				</div>

			</div>
		</div>
	</div>
	<div v-else>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h4> {{ this.errormsg }}</h4>
		</div>
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
