package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MathTestSuite struct {
	suite.Suite
}

func (suite *MathTestSuite) SetupSuite() {
	fmt.Println("在测试组运行之前总有些事情要做...")
}

func (suite *MathTestSuite) SetupTest() {
	fmt.Println("在测试运行之前总有些事情要做...")
}

func (suite *MathTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("要运行 %s 测试组里 %s 测试用例了...\n", suiteName, testName)
}

func (suite *MathTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("运行完 %s 测试组里 %s 测试用例了...\n", suiteName, testName)
}

func (suite *MathTestSuite) TearDownTest() {
	fmt.Println("在测试运行之后总有些事情要做...")
}

func (suite *MathTestSuite) TearDownSuite() {
	fmt.Println("在测试组运行之后总有些事情要做...")
}

func (suite *MathTestSuite) TestAdd() {
	suite.Equal(Add(1, 2), 3)
}

func (suite *MathTestSuite) TestSub() {
	suite.Equal(Sub(2, 1), 1)
}

func (suite *MathTestSuite) TestMul() {
	suite.Equal(Mul(9, 9), 81)
}

func (suite *MathTestSuite) TestDiv() {
	suite.Equal(Div(100, 10), 10)
}

func TestMathTestSuite(t *testing.T) {
	suite.Run(t, new(MathTestSuite))
}
