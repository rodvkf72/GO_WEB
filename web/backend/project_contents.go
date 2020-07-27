package backend

import (

)

//게시판 기능에 DB를 연결하였으므로 프로젝트 부분은 데이터를 직접 입력 및 수정하도록 구현 -> 코드의 중복이 많이 일어남
func return_contents(no string) []project_view {
	Project_View := project_view{}
	Project_Views := []project_view{}

	switch no {
	case "1":
		Project_View.Title = "공공데이터 파싱"
		Project_View.Content = "공공데이터 파싱을 활용하였다." 
		Project_View.Root = "/static/image/main_2.png"

	case "2":
		Project_View.Title = "캐릭터 달리기 게임"
		Project_View.Content = "MFC의 비트맵을 활용하여 캐릭터와 맵을 구현하여 게임을 제작하였다." 
		Project_View.Root = "/static/image/main_2.png"

	case "3":
		Project_View.Title = "인사관리 프로그램"
		Project_View.Content = "디자인 패턴을 활용하여 인사관리 프로그램을 제작하였다.\n" + 
							   "그 중 메멘토 패턴을 활용하여 이전 데이터로 되돌리는 부분을 맡았다." 
		Project_View.Root = "/static/image/main_2.png"

	case "4":
		Project_View.Title = "웹 사이트 제작"
		Project_View.Content = "Tomcat 서버를 활용하여 웹 사이트를 제작하였다.\n" +
							   "자유주제라 장르 통합형 웹 사이트로 공부의 기능과 오락의 기능을\n" +
							   "선택하여 사용할 수 있게 하는 사이트를 제작하였다."
		Project_View.Root = "/static/image/main_2.png"
		
	case "5":
		Project_View.Title = "자율주행 프로그램"
		Project_View.Content = "ROS(Robot Operating System)를 활용하여 터틀봇 자율주행 프로그램을 제작하였다.\n" +
							   "미로찾기와 주행 시험장에 대한 자율주행 프로그램을 제작하였으며 미로찾기에서는\n" +
							   "경로탐색과 백트래킹을 담당하고 자율주행에서는 중앙선 인식과 터틀봇의 회전,\n" +
							   "정지선 인식을 담당하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "6":
		Project_View.Title = "서비스센터 프로그램"
		Project_View.Content = "데이터베이스를 활용하여 서비스센터 프로그램을 제작하였다.\n" +
							   "기능은 서비스 접수하기, 서비스 접수 내역, 직원 관리, 고객 관리,\n" +
							   "부품관리, 이벤트, 이달의 직원이 있으며 담당했던 파트를 나누기에는\n" +
							   "다른 테이블과의 join, select, update 문의 반복이었기에 모든 팀원이\n" +
							   "모든 기능 부분에 참여하였다고 생각된다."
		Project_View.Root = "/static/image/main_2.png"

	case "7":
		Project_View.Title = "자동 차량 결제 애플리케이션"
		Project_View.Content = "OpenCV를 활용하여 자동 차량 결제 애플리케이션을 개발하였다.\n" +
							   "맥도날드나 스타벅스의 드라이브 쓰루 처럼 차량이 오면 직접 카드를\n" +
							   "건네주어 결제하지 않고 차량의 번호판을 OpenCV를 사용하여 인식하고\n" +
							   "구매 제품을 등록하여 차량과 연동된 카드에서 결제 되도록 하는 방식이었으나\n" +
							   "카드결제 모듈 발급에 있어서 사업자 등록증이 없는 관계로 모듈을 사용하지\n" +
							   "못하였으며 학습이 되어있지 않아 번호판 인식이 제대로 이루어지지 않았다.\n" +
							   "제대로 구현하려면 딥러닝을 통해 번호판 인식률을 높이고 사업자 등록을 통해\n" +
							   "결제모듈을 받거나 기존 사업자와의 제휴가 필요할 것으로 생각된다.\n" +
							   "담당 파트는 로그인, 회원가입, OpenCV로 스캔된 번호판의 데이터 전송 및\n" +
							   "결제모듈 창 띄우기 등 기능 전체를 담당하였다.\n" +
							   "진짜 OpenCV, 디자인 빼고는 혼자서 다 짰다. 심지어 OpenCV는 그냥 모듈..."
		Project_View.Root = "/static/image/main_2.png"

	case "8":
		Project_View.Title = "AR 구현동화"
		Project_View.Content = "유니티 뷰포리아의 AR 기능을 활용하여 구현동화 설계를 계획하였다.\n" +
							   "어린이들의 성장 발달에는 눈으로 보고 직접 학습하는 것이 도움이 되고\n" +
							   "휴대전화의 사용이 눈에 띄게 증가함에 따라 AR 구현동화 설계를 계획하였다.\n" +
							   "이는 보고서 양식의 캡스톤 디자인이므로 담당 파트는 없었으나 구현 가능성을\n" +
							   "보이기 위해 유니티의 뷰포리아를 사용해서 배경화면에 AR을 인식시키는 작업을 담당하였다"
		Project_View.Root = "/static/image/main_2.png"

	case "9":
		Project_View.Title = "건축정보 모바일 서비스 콘텐츠"
		Project_View.Content = "서울시 마포구에서 진행하는 건축정보 모방리 서비스 콘텐츠 개발 공모전에 참가하였다.\n" +
							   "공무에 도움이 되는 기능 구현도 포함이 되어서 건축물의 등급이 햇수와 직접적인 관찰로\n" +
							   "이루어지는데 등록된 건축물 마저도 여러 요소가 빠져있는 경우가 있어서 이를 한번에\n" +
							   "해결하기 위해 모바일로 해당 건축물에 대한 정보를 입력할 수 있게 하는 애플리케이션을\n" +
							   "개발하였다. 예선 광탈...\n" +
							   "맡은 부분은.. 혼자 다 했다. 2인 팀이었는데 건축과라서 애플리케이션 만들줄도 몰랐다.."
		Project_View.Root = "/static/image/main_2.png"

	case "10":
		Project_View.Title = "메일 시스템 유지/보수"
		Project_View.Content = "JSP를 활용하여 웹 메일 시스템에 대한 유지/보수 작업을 하였다.\n" +
							   "유지/보수는 JavaBeans 등을 활용한 코드 간소화, 파일 다중 첨부 기능 추가 등\n" +
							   "예방 유지/보수, 완전화 유지/보수, 적응 유지/보수, 수정 유지/보수를 하였다\n" +
							   "담당 파트는 파일 다중 첨부(완전화 유지/보수), XSS 공격 방지를 위한 필터링(예방 유지/보수),\n" +
							   "임시 보관함(완전화 유지/보수), 반응형 웹(적응 유지/보수), 오류코드 수정(수정 유지/보수) 이다."
		Project_View.Root = "/static/image/main_2.png"

	case "11":
		Project_View.Title = "여행 애플리케이션"
		Project_View.Content = "구글 맵을 활용하여 여행 애플리케이션을 제작하였다.\n" +
							   "기존의 여행 애플리케이션과의 차이점을 두기 위해 데이터베이스에 나이대별로\n" +
							   "해당 지역에 대한 선호율을 집계하고 이를 애플리케이션으로 보여준다. 따라서\n" +
							   "20대의 선호지역과 30대의 선호지역이 다르게 나타날 수 있으며 다모임이라는\n" +
							   "게시판을 이용하여 같이 여행 갈 인원을 모집할 수도 있다. 단, 게시판의 경우는\n" +
							   "악용될 수 있기 때문에 관라자의 2차 인증을 받아야 사용할 수 있다. 이는 대학생\n" +
							   "애플리케이션인 에브리타임에서도 사용하고 있는 방식이다.\n" +
							   "여기서 담당 파트는 로그인, 회원가입 기능과 다모임 게시판 구현이다." 
		Project_View.Root = "/static/image/main_2.png"

	case "12":
		Project_View.Title = "블로그 제작"
		Project_View.Content = "Apache 서버와 데이터베이스를 활용하여 블로그를 제작하고 있다.\n" +
							   "Go 언어로 웹 페이지 제작이 가능하다는 것을 보고 시작하였다.\n" +
							   "아직 공부해가면서 만드는 부분이라 이런 저런 부분에 문제가 있을 수 있다.\n" +
							   "Github에 올려서 저장하고 있는데 개인적인 프로젝트와 이름, 사진이\n" +
							   "포함되어 있어 현재는 비공개 상태로 올리고 있다."
		Project_View.Root = "/static/image/main_2.png"
	}

	Project_Views = append(Project_Views, Project_View)
	return Project_Views
}