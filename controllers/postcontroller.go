package controllers

func CreatePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var posts entities.Posts
	json.NewDecoder(r.Body).Decode(&posts)
	database.Instance.Create(&posts)
	json.NewEncoder(w).Encode(posts)
}

func checkIfPostExists(postId string) bool{
	var posts entities.Posts
	database.Instance.First(&posts, postId)
	if posts.PostId ==0{
		return false
	}
	return true
}

func GetPostById(w http.ResponseWriter, r *http.Request){
	postId := mux.Vars(r)["id"]
	if checkIfPostExists(postId) ==false{
		json.NewEncode(w).Encode("Post Not Found!")
		return
	}
	var post entities.Posts
	database.Instance.First(&post, postId)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request){
	var posts []entities.Posts
	database.Instance.Find(&posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request){
	postId := mux.Vars(r)["id"]
	if checkIfPostExists(postId) == false{
		json.NewEncoder(w).Encode("Post Not Found!")
		return
	}
	var post entities.Posts
	database.Instance.First(&post, postId)
	json.NewDecoder(r.Body).Decode(&post)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	postId := mux.Vars(r)["id"]
	if checkIfPostExists(postId)==false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Post Not Found!")
		return
	}
	var post entities.Posts
	database.Instance.Delete(&post, postId)
	json.NewEncoder(w).Encode("Post Deleted!")
}