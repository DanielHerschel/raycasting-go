package prefabs

type IDamageable interface {
	TakeDamage(amount int) bool
	IsAlive() bool
}
