package backend

/*
게시판 기능에 DB를 연결하였으므로 프로젝트 부분은 데이터를 직접 입력 및 수정하도록 구현 -> 코드의 중복이 많이 일어남
공부하는 입장에서 여러 방식을 적용, 테스트 해 보고 싶어서 하드코딩 한 부분
*/
func return_contents(no string) []project_view {
	Project_View := project_view{}
	Project_Views := []project_view{}

	switch no {
	case "1":
		Project_View.Title = "GO WEB"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - GoLand, Visual Studio Code, phpmyadmin(MySQL), Docker, Github\n" +
			"Language - Go, SQL, html, css, javascript\n\n" +
			"Go 언어로 웹 페이지 제작이 가능하다는 것을 보고 시작한 프로젝트\n" +
			"공부해가면서 만드는 것이라 이런 저런 부분에 문제가 있을 수 있다.\n\n"
		Project_View.Root = "/static/image/main_2.png"

	case "2":
		Project_View.Title = "Makers"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - Android Studio, Visual Studio Code, phpmyadmin(MySQL), Unity, Github\n" +
			"Language - Go, Java, SQL\n\n" +
			"기존의 여행 애플리케이션과의 차이점을 두기 위해 DB에 나이대별로 해당지역 선호율을\n" +
			"집계하고 이를 애플리케이션으로 보여준다. 따라서 20대의 선호지역과 30대의 선호지역이\n" +
			"다르게 나타날 수 있으며 다모임이라는 게시판을 이용하여 같이 여행 갈 인원을 모집할\n" +
			"수 있다. 단, 게시판의 경우 악용 방지를 위해 관리자의 2차 인증을 받아야 한다.\n" +
			"여기서 담당 파트는 로그인, 회원가입 기능과 다모임 게시판 구현이다."
		Project_View.Root = "/static/image/main_2.png"

	case "3":
		Project_View.Title = "Web Mail System"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - NetBeans IDE, James Server, Github, Notion\n" +
			"Language - Java, jsp, javascript\n\n" +
			"주어진 웹 메일 시스템에 대한 유지/보수 프로젝트\n" +
			"담당 파트는 파일 다중 첨부(완전화 유지/보수), 파일 첨부 쓰레드 처리, XSS 공격 방지를 위한 필터링(예방 유지/보수),\n" +
			"임시 보관함(완전화 유지/보수), 반응형 웹(적응 유지/보수), 오류코드 수정(수정 유지/보수) 이다."
		Project_View.Root = "/static/image/main_2.png"

	case "4":
		Project_View.Title = "Car Pay"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - Android Studio, Visual Studio Code, phpmyadmin(MySQL), Github\n" +
			"Language - Java, php, SQL\n\n" +
			"OpenCV를 활용하여 자동 차량 결제 애플리케이션을 개발하였다.\n" +
			"맥도날드나 스타벅스의 드라이브 쓰루 처럼 차량이 오면 직접 카드를\n" +
			"건네주어 결제하지 않고 차량의 번호판을 OpenCV를 사용하여 인식하고\n" +
			"구매 제품을 등록하여 차량과 연동된 카드에서 결제 되도록 하는 방식\n" +
			"그러나 카드결제 모듈 발급에 있어서 사업자 등록증이 없는 관계로 모듈을 사용하지\n" +
			"못하였으며 학습이 되어있지 않아 번호판 인식이 제대로 이루어지지 않았다.\n" +
			"제대로 구현하려면 딥러닝을 통해 번호판 인식률을 높이고 사업자 등록을 통해\n" +
			"결제모듈을 받거나 기존 사업자와의 제휴가 필요할 것으로 생각된다.\n" +
			"담당 파트는 로그인, 회원가입, OpenCV로 스캔된 번호판의 데이터 전송 및\n" +
			"결제모듈 창 띄우기 등 기능 전체를 담당하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "5":
		Project_View.Title = "건축물음표"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - Android Studio, Visual Studio Code, phpmyadmin(MySQL)\n" +
			"Language - php, Java, SQL\n\n" +
			"서울시 마포구에서 진행하는 건축정보 모바일 서비스 콘텐츠 개발 공모전에 참가하였다.\n" +
			"공무에 도움이 되는 기능 구현도 포함이 되어서 건축물의 등급이 햇수와 직접적인 관찰로\n" +
			"이루어지는데 등록된 건축물 마저도 여러 요소가 빠져있는 경우가 있어서 이를 한번에\n" +
			"해결하기 위해 모바일로 해당 건축물에 대한 정보를 입력할 수 있게 하는 애플리케이션을\n" +
			"개발하였다. 예선 광탈...\n" +
			"실질적으로는 1인 개발이었다. 다른 팀원은 건축과라 프로그래밍을 몰라서 정보 제공을 하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "6":
		Project_View.Title = "터틀봇 자율주행 프로그램"
		Project_View.Content = "OS - ubuntu 16.04, ROS\n" +
			"Tool - PyCharm, Gazebo, Rviz, Github\n" +
			"Language - python\n\n" +
			"ROS(Robot Operating System)를 활용한 터틀봇 자율주행 프로젝트\n" +
			"미로찾기와 주행 시험장에 대한 자율주행 프로그램을 제작하였으며 미로찾기에서는\n" +
			"경로탐색과 백트래킹을 담당하고 자율주행에서는 중앙선 인식과 터틀봇의 회전,\n" +
			"정지선 인식을 담당하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "7":
		Project_View.Title = "서비스센터 프로그램"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - Visual Studio, Github\n" +
			"Language - Pro C, SQL\n\n" +
			"DB를 활용한 서비스센터 프로그램 제작 프로젝트\n" +
			"기능은 서비스 접수하기, 서비스 접수 내역, 직원 관리, 고객 관리,\n" +
			"부품관리, 이벤트, 이달의 직원이 있으며 담당했던 파트를 나누기에는\n" +
			"다른 테이블과의 join, select, update 문의 반복이었기에 모든 팀원이\n" +
			"모든 기능 부분에 참여하였다고 생각된다."
		Project_View.Root = "/static/image/main_2.png"

	case "8":
		Project_View.Title = "인사관리 프로그램"
		Project_View.Content = "OS - Window 10\n" +
			"Tool - Eclipse, BOUML" +
			"디자인 패턴을 활용한 인사관리 프로그램 제작 프로젝트\n" +
			"그 중 메멘토 패턴을 활용하여 이전 데이터로 되돌리는 부분을 맡았다."
		Project_View.Root = "/static/image/main_2.png"

	case "9":
		Project_View.Title = "웹 사이트 제작"
		Project_View.Content = "Tomcat 서버를 활용하여 웹 사이트를 제작하였다.\n" +
			"자유주제라 장르 통합형 웹 사이트로 공부의 기능과 오락의 기능을\n" +
			"선택하여 사용할 수 있게 하는 사이트를 제작하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "10":
		Project_View.Title = "AR 구현동화"
		Project_View.Content = "유니티 뷰포리아의 AR 기능을 활용하여 구현동화 설계를 계획하였다.\n" +
			"어린이들의 성장 발달에는 눈으로 보고 직접 학습하는 것이 도움이 되고\n" +
			"휴대전화의 사용이 눈에 띄게 증가함에 따라 AR 구현동화 설계를 계획하였다.\n" +
			"이는 보고서 양식의 캡스톤 디자인이므로 담당 파트는 없었으나 구현 가능성을\n" +
			"보이기 위해 유니티의 뷰포리아를 사용해서 배경화면에 AR을 인식시키는 작업을 담당하였다"
		Project_View.Root = "/static/image/main_2.png"

	case "11":
		Project_View.Title = "공공데이터 파싱"
		Project_View.Content = "공공데이터 파싱을 활용하였다."
		Project_View.Root = "/static/image/main_2.png"

	case "12":
		Project_View.Title = "캐릭터 달리기 게임"
		Project_View.Content = "MFC의 비트맵을 활용하여 캐릭터와 맵을 구현하여 게임을 제작하였다."
		Project_View.Root = "/static/image/main_2.png"

	}

	Project_Views = append(Project_Views, Project_View)
	return Project_Views
}
