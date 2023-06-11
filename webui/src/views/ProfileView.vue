<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			photos: [],
			comment: [''],
			usernameToSearch: "",
			listFollower: [],
			listFollowing: [],
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
				if (e.response && e.response.status === 404) {
					this.errormsg = "404";
				} else {
					this.errormsg = e.toString();
				}
			}
			console.log("Prendo l'user")

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
			console.log("Prendo le foto")

			try {
				let response = await this.$axios.get("/users/"+ this.user.id +"/following/", {
					headers: {
						Authorization: this.user.id
					}
				});
				this.listFollowing = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			console.log("Prendo i following")

			try {
				let response = await this.$axios.get("/users/"+ this.user.id +"/follower/", {
					headers: {
						Authorization: this.user.id
					}
				});
				this.listFollower = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			console.log("Prendo i follower")

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
			console.log("Controllo il bottone")
			console.log(this.followed)

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

			console.log("Controllo il secondo bottone")
			console.log(this.banned)
			console.log(this.photos)
			this.loading = false;
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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.followed = true
			this.loading = false;
			this.refresh();
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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.followed = false
			this.loading = false;
			this.refresh();
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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.banned = true
			this.loading = false;
			this.refresh();
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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.banned = false
			this.loading = false;
			this.refresh()
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
				this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async showPhoto(photoid){
			this.$router.push("/photo/"+ photoid);
			this.refresh()
		},

		async changeUsername() {
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
			<h1 class="h2" v-else>My profile</h1>

			<div class="btn-group">
				<button type="button" class="btn btn-default">Following: {{this.user.following}}</button>
				<button type="button" class="btn btn-default dropdown-toggle dropdown-toggle-split" data-bs-toggle="dropdown" data-bs-reference="parent">
					<span class="visually-hidden">Toggle Dropdown</span>
				</button>
				<div class="dropdown-menu">
					<a class="dropdown-item" v-for="us in this.listFollowing" :key="us">{{ us.Username}}</a>
				</div>
			</div>

			<div class="btn-group">
				<button type="button" class="btn btn-default">Follower: {{this.user.follower}}</button>
				<button type="button" class="btn btn-default dropdown-toggle dropdown-toggle-split" data-bs-toggle="dropdown" data-bs-reference="parent">
					<span class="visually-hidden">Toggle Dropdown</span>
				</button>
				<div class="dropdown-menu">
					<a class="dropdown-item" v-for="use in this.listFollower" :key="use">{{ use.Username}}</a>
				</div>
			</div>


			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="input-group me-2" v-if="this.user.id == this.loggedId">
					<input type="string" class="form-control" v-model="newUsername" placeholder="New username">
					<button type="button" class="btn btn-primary" @click="changeUsername()" >
						<svg class="feather"> <use href="/feather-sprite-v4.29.0.svg#pen-tool"></use> </svg>
					</button>
				</div>
				<div v-else>
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

		<div class="card" v-if="this.photos == null">
			<div class="card-header">
				<p>{{this.user.username}} hasn't upload any photos yet</p>
			</div>
		</div>
		<div class="card" v-else>
			<div class="card-header">
				Photos
			</div>
			<div class="card-body">
				<div class="container" ></div>
					<div class="row" v-if="!loading">
						<div class="col-md-6" v-for="p in photos" :key="p">
							<div class="card-text">
								<div @click="showPhoto(p.ID)">
									<img :src="'data:image;base64,' + p.Picture" style="height: 250px;"/><br />
								</div>
								User: {{ p.User_id }}<br />
								Date: {{ p.Date_time }}<br />
								Likes: {{ p.Likes }}<br />
								Comments: {{ p.Comments }}<br />
							</div>
							<div class="input-group mb-3">
								<input type="string" class="form-control" v-model="comment[p.ID]" placeholder="Write your comment here">
								<button type="button" :disabled="this.comment[p.ID] == null " class="btn btn-primary" @click="commentPhoto(p.ID, this.comment[p.ID])" >
									<svg class="feather"> <use href="/feather-sprite-v4.29.0.svg#send"></use> </svg>
								</button>
							</div>
						</div>
					</div>
				
			</div>
		</div>
	</div>
	<div v-else>
		<div v-if="this.errormsg == '404' ">
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h4> The user {{ this.$route.params.username }} doesn't not exist </h4>
			</div>
			<div>
				<h6> Try with a different username</h6>
			</div>
		</div>
		<div v-else>
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h4> {{ this.errormsg }}</h4>
			</div>
		</div>
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
