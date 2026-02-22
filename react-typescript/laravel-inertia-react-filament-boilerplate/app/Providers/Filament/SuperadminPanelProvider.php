<?php

namespace App\Providers\Filament;

use Filament\Http\Middleware\Authenticate;
use Filament\Http\Middleware\AuthenticateSession;
use Filament\Http\Middleware\DisableBladeIconComponents;
use Filament\Http\Middleware\DispatchServingFilamentEvent;
use Filament\Pages;
use Filament\Panel;
use Filament\PanelProvider;
use Filament\Widgets;
use Illuminate\Cookie\Middleware\AddQueuedCookiesToResponse;
use Illuminate\Cookie\Middleware\EncryptCookies;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use Illuminate\Routing\Middleware\SubstituteBindings;
use Illuminate\Session\Middleware\StartSession;
use Illuminate\View\Middleware\ShareErrorsFromSession;

class SuperadminPanelProvider extends PanelProvider
{
    public function panel(Panel $panel): Panel
    {
        return $panel
            ->default()
            ->id('superadmin')
            ->path('superadmin')
            ->login()
            ->brandName('Portofolio')
            ->favicon(asset('favicon.ico'))
            ->colors([
                'primary' => [
                    50  => '#FFF0FA',
                    100 => '#FFD6F3',
                    200 => '#FFB3E9',
                    300 => '#FF80D9',
                    400 => '#FF40C0',
                    500 => '#FF00AA',
                    600 => '#CC0088',
                    700 => '#990066',
                    800 => '#660044',
                    900 => '#330022',
                    950 => '#1A0011',
                ],
                'gray' => [
                    50  => '#E0FFFE',
                    100 => '#B3FFFC',
                    200 => '#80FFF9',
                    300 => '#4DFFF6',
                    400 => '#1AFFF3',
                    500 => '#00FFF1',
                    600 => '#00CCB8',
                    700 => '#009988',
                    800 => '#111827',
                    900 => '#0D1117',
                    950 => '#080D12',
                ],
            ])
            ->font('Manrope')
            ->darkMode(true, isForced: true)
            ->sidebarCollapsibleOnDesktop()
            ->maxContentWidth('full')
            ->discoverResources(in: app_path('Filament/Resources'), for: 'App\\Filament\\Resources')
            ->discoverPages(in: app_path('Filament/Pages'), for: 'App\\Filament\\Pages')
            ->pages([
                Pages\Dashboard::class,
            ])
            ->discoverWidgets(in: app_path('Filament/Widgets'), for: 'App\\Filament\\Widgets')
            ->widgets([
                Widgets\AccountWidget::class,
            ])
            ->middleware([
                EncryptCookies::class,
                AddQueuedCookiesToResponse::class,
                StartSession::class,
                AuthenticateSession::class,
                ShareErrorsFromSession::class,
                VerifyCsrfToken::class,
                SubstituteBindings::class,
                DisableBladeIconComponents::class,
                DispatchServingFilamentEvent::class,
            ])
            ->authMiddleware([
                Authenticate::class,
            ]);
    }
}
