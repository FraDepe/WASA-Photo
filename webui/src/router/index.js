import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue' 
import HomeView from '../views/HomeView.vue'
import UploadPhotoView from '../views/UploadPhotoView.vue'
import ProfileView from '../views/ProfileView.vue'
import PhotoView from '../views/PhotoView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/new', component: UploadPhotoView},
		{path: '/user/:username', component: ProfileView},
		{path: '/photo/:photoid', component: PhotoView},
	]
})

export default router
