<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			photos: [],
			comment: [''],
			usernameToSearch: "",
			followed: false,
			follower: {
				followerid: 0,
				followedid: 0
			},
			banned: false,
			ban: {
				userid: 0,
				bannedid: 0
			},
			user: {
				id: 0,
				username: 0,
				follower: 0,
				following: 0,
				banned: 0,
				photos: 0
			},
			loggedId: 0
            
		}
	},
	methods: {

		async refresh() {
			this.loading = true;
			this.errormsg = null;

			try {
				let response = await this.$axios.get("/users/"+ localStorage.userid +"/profile/" + localStorage.usernameToSearch, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.user = response.data;
				this.loggedId = localStorage.userid
			} catch (e) {
				this.errormsg = e.toString();
			}

			console.log(this.errormsg)

			try {
				let response = await this.$axios.get("/users/"+ localStorage.userid +"/profile/" + localStorage.usernameToSearch + "/", {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.photos = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			console.log(this.errormsg)

			try {
				let response = await this.$axios.get("/users/"+ localStorage.userid +"/following/" + this.user.id, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.follower = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			if (this.follower.followedid == this.user.id) {
				this.followed = true
			}
			else {
				this.followed = false
			}

			console.log(this.errormsg)

			try {
				let response = await this.$axios.get("/users/"+ localStorage.userid +"/banned/" + this.user.id);
				this.ban = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			if (this.ban.bannedid == this.user.id) {
				this.banned = true
			}
			else {
				this.banned = false
			}

			console.log(this.errormsg)

			this.loading = false;
			console.log(this.user.id)
			console.log(this.loggedId)
		},

		async followUser() {
			this.loading = true;
			this.errormsg = null;

			try {
				await this.$axios.put("/users/" + localStorage.userid + "/following/" + this.user.id, {}, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.following = true
				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}

			this.loading = false;
		},

		async unfollowUser() {
			this.loading = true;
			this.errormsg = null;

			try {
				await this.$axios.delete("/users/" + localStorage.userid + "/following/" + this.user.id, {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.following = false
				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async banUser() {
			this.loading = true;
			this.errormsg = null;

			try {
				await this.$axios.put("/users/" + localStorage.userid + "/banned/" + this.user.id, {}, {
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

		async unbanUser() {
			this.loading = true;
			this.errormsg = null;

			try {
				await this.$axios.delete("/users/" + localStorage.userid + "/banned/" + this.user.id, {
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

		async commentPhoto(photoid, stringa) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.post("/photos/" + photoid + "/comments/", stringa , {
					headers: {
						Authorization: localStorage.userid
					}
				});
				this.comment = ['']
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

		async showPhoto(photoid){
			this.$router.push("/photo/"+ photoid);
			this.refresh()
		},

		async changeUsername(string) {
			await this.$axios.put("/users/" + localStorage.userid + "/changeUsername",this.newUsername , {
					headers: {
						Authorization: localStorage.userid
					}
				});
			localStorage.usernameToSearch = this.newUsername
			this.$router.replace("/user/"+ this.newUsername);
			this.newUsername = ""
			this.refresh()
		},
	},

	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div v-if="this.errormsg == null">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2" v-if="this.user.id != this.loggedId">Profile of {{this.user.username}}</h1>
			<h1 class="h2" v-else>Your stream</h1>
			<h6 class="label"> Follower: {{this.user.follower}}</h6>
			<h6 class="label"> Following: {{this.user.following}}</h6>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="input-group me-2" v-if="this.user.id == this.loggedId">
					<input type="string" class="form-control" v-model="newUsername" placeholder="New username">
					<button type="button" class="btn btn-primary" @click="changeUsername(this.newUsername)" >
						Change
					</button>
				</div>
				<div class="input-group me-2">
					<input type="string" class="form-control" v-model="usernameToSearch" placeholder="Search here for a user">
					<button type="button" class="btn btn-primary" @click="getUserProfile(this.usernameToSearch)" >
						Search
					</button>
				</div>
					<div v-if="this.user.id != this.loggedId">
						<div class="btn-group me-2" v-if="this.followed == false"> 
							<button type="button" class="btn btn-sm btn-success" @click="followUser">
								Follow
							</button>
						</div>
						<div class="btn-group me-2" v-else>
							<button type="button" class="btn btn-sm btn-danger" @click="unfollowUser">
								Unfollow
							</button>
						</div>
						<div class="btn-group me-2" v-if="this.banned == false"> 
							<button type="button" class="btn btn-sm btn-outline-danger" @click="banUser">
								Ban
							</button>
						</div>
						<div class="btn-group me-2" v-else>
							<button type="button" class="btn btn-sm btn-outline-danger" @click="unbanUser">
								Unban
							</button>
						</div>
					</div>
				</div>
			</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="photos.length === 0">
			<div class="card-body">
				<p>{{this.user.username}} hasn't upload any photos yet</p>
			</div>
		</div>
		<div class="card" v-else>
			<div class="card-header">
				Photos
			</div>
			<div class="card-body">
				<div class="container" ></div>
					<div class="row" >
						<div class="col-md-6" v-if="!loading" v-for="p in photos">
							<p class="card-text">
								<div @click="showPhoto(p.ID)">
									<img :src="'data:image;base64,' + p.Picture" style="height: 250px;"/><br />
								</div>
								User: {{ p.User_id }}<br />
								Date: {{ p.Date_time }}<br />
								Likes: {{ p.Likes }}<br />
								Comments: {{ p.Comments }}<br />
							</p>
							<div class="input-group mb-3">
								<input type="string" class="form-control" v-model="comment[p.ID]" placeholder="Write your comment here">
								<button type="button" :disabled="this.comment[p.ID] == null " class="btn btn-primary" @click="commentPhoto(p.ID, this.comment[p.ID])" >Post</button>
							</div>
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
