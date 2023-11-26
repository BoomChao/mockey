package mockey

import (
	"fmt"
	"testing"

	"github.com/BoomChao/mockey/ani"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func Foo(in string) string {
	return in
}

type A struct{}

func (a A) Foo(in string) string { return in }

var Bar = 0

func TestMockXXX(t *testing.T) {
	PatchConvey("TestMockXXX", t, func() {
		Mock(Foo).Return("c").Build()   // mock函数
		Mock(A.Foo).Return("c").Build() // mock方法
		MockValue(&Bar).To(1)           // mock变量

		So(Foo("a"), ShouldEqual, "c")        // 断言`Foo`成功mock
		So(new(A).Foo("b"), ShouldEqual, "c") // 断言`A.Foo`成功mock
		So(Bar, ShouldEqual, 1)               // 断言`Bar`成功mock
	})
	// `PatchConvey`外自动释放mock
	fmt.Println(Foo("a"))        // a
	fmt.Println(new(A).Foo("b")) // b
	fmt.Println(Bar)             // 0
}

func TestInterface(t *testing.T) {
	PatchConvey("TestInterface", t, func() {
		PatchConvey("interface-dog", func() {
			result := "ww"
			// 对接口成员变量的方法进行mock,这里dog是可导出的类型的，所以直接可以引用,但是如果
			// 我们想要用cat来实例化这个animal，则就不能用这个方法了
			dog := ani.Dog{}
			Mock(GetMethod(dog, "Speak")).Return(result).Build()
			zoo := &ani.Zoo{Ani: dog}
			got := zoo.AniSpeak("sss")
			So(got, ShouldEqual, result)
		})

		// 未导出类型直接调用NewStruct方法来进行构造
		PatchConvey("interface-cat", func() {
			result := "ww"
			cat := ani.NewCat()
			Mock(GetMethod(cat, "Speak")).Return(result).Build()
			zoo := &ani.Zoo{Ani: cat}
			got := zoo.AniSpeak("sss")
			So(got, ShouldEqual, result)
		})
	})
}
