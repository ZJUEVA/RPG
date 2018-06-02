package scene

type Scene interface {
	Enter(manager *SceneManager)
	Update()
	Leave()
}
